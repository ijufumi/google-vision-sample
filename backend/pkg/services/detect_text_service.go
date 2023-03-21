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
	jobFileRepository repositories.JobFileRepository,
	imageConversionService ImageConversionService,
	logger *zap.Logger,
	db *gorm.DB,
) DetectTextService {
	return &detectTextService{
		storageAPIClient:        storageAPIClient,
		visionAPIClient:         visionAPIClient,
		jobRepository:           jobRepository,
		extractedTextRepository: extractedTextRepository,
		jobFileRepository:       jobFileRepository,
		imageConversionService:  imageConversionService,
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

	fileInfo, _ := file.Stat()
	splitFileName := strings.Split(file.Name(), "/")
	err = s.db.Transaction(func(tx *gorm.DB) error {
		job := &entities.Job{
			ID:     id,
			Status: enums.JobStatus_Runing,
		}
		err = s.jobRepository.Create(tx, job)
		if err != nil {
			return err
		}
		err = s.jobFileRepository.Create(tx, &entities.JobFile{
			ID:          inputFileID,
			JobID:       id,
			IsOutput:    false,
			FileKey:     key,
			FileName:    splitFileName[len(splitFileName)-1],
			ContentType: contentType,
			Size:        fileInfo.Size(),
		})
		return err
	})

	if err != nil {
		return err
	}

	go func() {
		err := s.processDetectText(id, key, file.Name())
		if err != nil {
			s.logger.Error(fmt.Sprintf("%v was occurred.", err))
		}
	}()
	return nil
}

func (s *detectTextService) processDetectText(id, key, imageFilePath string) error {
	job, err := s.jobRepository.GetByID(s.db, id)
	if err != nil {
		return err
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		outputKey, err := s.visionAPIClient.DetectText(key)
		if err != nil {
			return err
		}

		queryFiles, err := s.storageAPIClient.QueryFiles(outputKey)
		if err != nil {
			return err
		}

		if len(queryFiles) == 0 {
			return errors.New("output does not exist")
		}

		outputFile, err := s.storageAPIClient.DownloadFile(queryFiles[0])
		if err != nil {
			return err
		}

		fileStat, _ := outputFile.Stat()
		outputFileKey := queryFiles[0]
		splitOutputFileKey := strings.Split(outputFileKey, "/")
		err = s.jobFileRepository.Create(tx, &entities.JobFile{
			ID:          utils.NewULID(),
			JobID:       id,
			IsOutput:    true,
			FileKey:     outputFileKey,
			FileName:    splitOutputFileKey[len(splitOutputFileKey)-1],
			ContentType: "application/json",
			Size:        fileStat.Size(),
		})
		if err != nil {
			return err
		}

		bytes, err := io.ReadAll(outputFile)
		if err != nil {
			return err
		}

		var detectResponse googleModels.DetectTextResponses
		if err = json.Unmarshal(bytes, &detectResponse); err != nil {
			return err
		}
		extractedTexts := make([]*entities.ExtractedText, 0)
		orientation, err := s.imageConversionService.DetectOrientation(imageFilePath)
		if err != nil {
			return err
		}
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
						points := s.imageConversionService.ConvertPoints(vertices, orientation)
						xArray := make([]float64, 0)
						yArray := make([]float64, 0)
						for _, point := range points {
							xArray = append(xArray, point[0])
							yArray = append(yArray, point[1])
						}
						top := utils.MinInArray(yArray...)
						bottom := utils.MaxInArray(yArray...)
						left := utils.MinInArray(xArray...)
						right := utils.MaxInArray(xArray...)
						extractedText := &entities.ExtractedText{
							ID:     utils.NewULID(),
							JobID:  id,
							Text:   texts,
							Top:    top,
							Bottom: bottom,
							Left:   left,
							Right:  right,
						}

						extractedTexts = append(extractedTexts, extractedText)
					}
				}
			}
		}
		return s.extractedTextRepository.Create(tx, extractedTexts...)
	})
	status := enums.JobStatus_Succeeded
	if err != nil {
		status = enums.JobStatus_Failed
	}
	job.Status = status

	return s.jobRepository.Update(s.db, job)
}

func (s *detectTextService) DeleteResult(id string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		err := s.extractedTextRepository.DeleteByExtractionResultID(tx, id)
		if err != nil {
			return err
		}
		files, err := s.jobFileRepository.GetByJobID(tx, id)
		if err != nil {
			return err
		}
		for _, file := range files {
			err = s.storageAPIClient.DeleteFile(file.FileKey)
			if err != nil {
				s.logger.Error(err.Error())
			}
		}
		err = s.jobFileRepository.DeleteByJobID(tx, id)
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
			ID:        extractedText.ID,
			JobID:     extractedText.JobID,
			Text:      extractedText.Text,
			Top:       extractedText.Top,
			Bottom:    extractedText.Bottom,
			Left:      extractedText.Left,
			Right:     extractedText.Right,
			CreatedAt: extractedText.CreatedAt.Unix(),
			UpdatedAt: extractedText.UpdatedAt.Unix(),
		})
	}
	for _, file := range entity.JobFiles {
		files = append(files, models.JobFile{
			ID:          file.ID,
			JobID:       file.JobID,
			IsOutput:    file.IsOutput,
			FileKey:     file.FileKey,
			FileName:    file.FileName,
			ContentType: file.ContentType,
			Size:        file.Size,
			CreatedAt:   file.CreatedAt.Unix(),
			UpdatedAt:   file.UpdatedAt.Unix(),
		})
	}
	return &models.Job{
		ID:             entity.ID,
		Status:         entity.Status,
		CreatedAt:      entity.CreatedAt.Unix(),
		UpdatedAt:      entity.UpdatedAt.Unix(),
		ExtractedTexts: extractedTexts,
		JobFiles:       files,
	}
}

type detectTextService struct {
	storageAPIClient        clients.StorageAPIClient
	visionAPIClient         clients.VisionAPIClient
	jobRepository           repositories.JobRepository
	extractedTextRepository repositories.ExtractedTextRepository
	jobFileRepository       repositories.JobFileRepository
	imageConversionService  ImageConversionService
	db                      *gorm.DB
	logger                  *zap.Logger
}
