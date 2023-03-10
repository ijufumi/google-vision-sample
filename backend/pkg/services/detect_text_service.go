package services

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/repositories"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/clients"
	googleModels "github.com/ijufumi/google-vision-sample/pkg/gateways/google/models"
	"github.com/ijufumi/google-vision-sample/pkg/models"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type DetectTextService interface {
	GetResults() ([]*models.Job, error)
	GetResultByID(id string) (*models.Job, error)
	GetSignedURL(key string) (*models.SignedURL, error)
	DetectTexts(file *os.File, contentType string) error
	DeleteResult(id string) error
}

func NewDetectTextService(
	storageAPIClient clients.StorageAPIClient,
	visionAPIClient clients.VisionAPIClient,
	jobRepository repositories.JobRepository,
	extractedTextRepository repositories.ExtractedTextRepository,
	fileRepository repositories.JobFileRepository,
	logger *zap.Logger,
	db *gorm.DB,
) DetectTextService {
	return &detectTextService{
		storageAPIClient:        storageAPIClient,
		visionAPIClient:         visionAPIClient,
		jobRepository:           jobRepository,
		extractedTextRepository: extractedTextRepository,
		fileRepository:          fileRepository,
		logger:                  logger,
		db:                      db,
	}
}

func (s *detectTextService) GetResults() ([]*models.Job, error) {
	results, err := s.jobRepository.GetAll(s.db)
	if err != nil {
		return nil, err
	}

	extractionResults := make([]*models.Job, 0)
	for _, result := range results {
		extractionResults = append(extractionResults, s.buildExtractionResultResponse(result))
	}
	return extractionResults, nil
}

func (s *detectTextService) GetResultByID(id string) (*models.Job, error) {
	result, err := s.jobRepository.GetByID(s.db, id)
	if err != nil {
		return nil, err
	}
	return s.buildExtractionResultResponse(result), nil
}

func (s *detectTextService) GetSignedURL(key string) (*models.SignedURL, error) {
	signedURL, err := s.storageAPIClient.SignedURL(key)
	if err != nil {
		return nil, err
	}
	return &models.SignedURL{URL: signedURL}, nil
}

func (s *detectTextService) DetectTexts(file *os.File, contentType string) error {
	id := utils.NewULID()
	inputFileID := utils.NewULID()
	key := fmt.Sprintf("%s/%s/%s", id, inputFileID, filepath.Base(file.Name()))

	err := s.storageAPIClient.UploadFile(key, file, contentType)
	if err != nil {
		return err
	}

	var result *entities.Job

	status := enums.JobStatus_Succeeded
	defer func() {
		if result != nil {
			result.Status = status
			_ = s.jobRepository.Update(s.db, result)
		}
	}()

	return s.db.Transaction(func(tx *gorm.DB) error {
		extractionResult := &entities.Job{
			ID:     id,
			Status: enums.JobStatus_Runing,
		}
		err = s.jobRepository.Create(tx, extractionResult)
		if err != nil {
			return err
		}
		err = s.fileRepository.Create(tx, &entities.JobFile{
			ID:                 inputFileID,
			ExtractionResultID: id,
			IsOutput:           false,
			FileKey:            key,
			FileName:           file.Name(),
			ContentType:        contentType,
			Size:               0,
		})
		if err != nil {
			return err
		}
		extractionResult, err = s.jobRepository.GetByID(tx, id)
		if err != nil {
			return err
		}
		result = extractionResult
		outputKey, err := s.visionAPIClient.DetectText(key)
		if err != nil {
			status = enums.JobStatus_Failed
			return err
		}

		queryFiles, err := s.storageAPIClient.QueryFiles(outputKey)
		if err != nil {
			status = enums.JobStatus_Failed
			return err
		}

		if len(queryFiles) == 0 {
			status = enums.JobStatus_Failed
			return errors.New("output does not exist")
		}

		outputFile, err := s.storageAPIClient.DownloadFile(queryFiles[0])
		if err != nil {
			status = enums.JobStatus_Failed
			return err
		}

		outputFileKey := queryFiles[0]
		splitOutputFileKey := strings.Split(outputFileKey, "/")
		err = s.fileRepository.Create(tx, &entities.JobFile{
			ID:                 utils.NewULID(),
			ExtractionResultID: id,
			IsOutput:           true,
			FileKey:            outputFileKey,
			FileName:           splitOutputFileKey[len(splitOutputFileKey)-1],
			ContentType:        "application/json",
			Size:               0,
		})
		if err != nil {
			status = enums.JobStatus_Failed
			return err
		}

		err = s.jobRepository.Update(tx, result)
		if err != nil {
			status = enums.JobStatus_Failed
			return err
		}
		extractionResult, err = s.jobRepository.GetByID(tx, id)
		if err != nil {
			status = enums.JobStatus_Failed
			return err
		}
		result = extractionResult

		bytes, err := io.ReadAll(outputFile)
		if err != nil {
			status = enums.JobStatus_Failed
			return err
		}

		var detectResponse googleModels.DetectTextResponses
		if err = json.Unmarshal(bytes, &detectResponse); err != nil {
			status = enums.JobStatus_Failed
			return err
		}

		extractedTexts := make([]*entities.ExtractedText, 0)
		for _, response := range detectResponse.Responses {
			for _, page := range response.FullTextAnnotation.Pages {
				for _, block := range page.Blocks {
					for _, paragraph := range block.Paragraphs {
						texts := ""
						for _, word := range paragraph.Words {
							for _, symbol := range word.Symbols {
								texts += symbol.Text
							}
						}
						vertices := paragraph.BoundingBox.Vertices
						xArray := make([]float64, 0)
						yArray := make([]float64, 0)
						for _, _vertices := range vertices {
							xArray = append(xArray, _vertices.X)
							yArray = append(yArray, _vertices.Y)
						}
						top := utils.MinInArray(xArray...)
						bottom := utils.MaxInArray(xArray...)
						left := utils.MinInArray(yArray...)
						right := utils.MaxInArray(yArray...)
						extractedText := &entities.ExtractedText{
							ID:                 utils.NewULID(),
							ExtractionResultID: extractionResult.ID,
							Text:               texts,
							Top:                top,
							Bottom:             bottom,
							Left:               left,
							Right:              right,
						}

						extractedTexts = append(extractedTexts, extractedText)
					}
				}
			}
		}
		return s.extractedTextRepository.Create(tx, extractedTexts...)
	})
}

func (s *detectTextService) DeleteResult(id string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		err := s.extractedTextRepository.DeleteByExtractionResultID(tx, id)
		if err != nil {
			return err
		}
		files, err := s.fileRepository.GetByExtractionResultID(tx, id)
		if err != nil {
			return err
		}
		for _, file := range files {
			err = s.storageAPIClient.DeleteFile(file.FileKey)
			if err != nil {
				s.logger.Error(err.Error())
			}
		}
		err = s.fileRepository.DeleteByExtractionResultID(tx, id)
		if err != nil {
			return err
		}

		return s.jobRepository.Delete(tx, id)
	})
}

func (s *detectTextService) buildExtractionResultResponse(entity *entities.Job) *models.Job {
	extractedTexts := make([]models.ExtractedText, 0)
	files := make([]models.JobFile, 0)

	for _, extractedText := range entity.ExtractedTexts {
		extractedTexts = append(extractedTexts, models.ExtractedText{
			ID:                 extractedText.ID,
			ExtractionResultID: extractedText.ExtractionResultID,
			Text:               extractedText.Text,
			Top:                extractedText.Top,
			Bottom:             extractedText.Bottom,
			Left:               extractedText.Left,
			Right:              extractedText.Right,
			CreatedAt:          extractedText.CreatedAt.Unix(),
			UpdatedAt:          extractedText.UpdatedAt.Unix(),
		})
	}
	for _, file := range entity.Files {
		files = append(files, models.JobFile{
			ID:                 file.ID,
			ExtractionResultID: file.ExtractionResultID,
			IsOutput:           file.IsOutput,
			FileKey:            file.FileKey,
			FileName:           file.FileName,
			ContentType:        file.ContentType,
			Size:               file.Size,
			CreatedAt:          file.CreatedAt.Unix(),
			UpdatedAt:          file.UpdatedAt.Unix(),
		})
	}
	return &models.Job{
		ID:             entity.ID,
		Status:         entity.Status,
		CreatedAt:      entity.CreatedAt.Unix(),
		UpdatedAt:      entity.UpdatedAt.Unix(),
		ExtractedTexts: extractedTexts,
		Files:          files,
	}
}

type detectTextService struct {
	storageAPIClient        clients.StorageAPIClient
	visionAPIClient         clients.VisionAPIClient
	jobRepository           repositories.JobRepository
	extractedTextRepository repositories.ExtractedTextRepository
	fileRepository          repositories.JobFileRepository
	db                      *gorm.DB
	logger                  *zap.Logger
}
