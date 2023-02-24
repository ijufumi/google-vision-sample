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
)

type DetectTextService interface {
	GetResults() ([]*models.ExtractionResult, error)
	GetResultByID(id string) (*models.ExtractionResult, error)
	DetectTexts(file *os.File) error
	DeleteResult(id string) error
}

func NewDetectTextService(
	storageAPIClient clients.StorageAPIClient,
	visionAPIClient clients.VisionAPIClient,
	extractionResultRepository repositories.ExtractionResultRepository,
	extractedTextRepository repositories.ExtractedTextRepository,
	logger *zap.Logger,
	db *gorm.DB,
) DetectTextService {
	return &detectTextService{
		storageAPIClient:           storageAPIClient,
		visionAPIClient:            visionAPIClient,
		extractionResultRepository: extractionResultRepository,
		extractedTextRepository:    extractedTextRepository,
		logger:                     logger,
		db:                         db,
	}
}

func (s *detectTextService) GetResults() ([]*models.ExtractionResult, error) {
	results, err := s.extractionResultRepository.GetAll(s.db)
	if err != nil {
		return nil, err
	}

	extractionResults := make([]*models.ExtractionResult, 0)
	for _, result := range results {
		extractionResults = append(extractionResults, s.buildExtractionResultResponse(result))
	}
	return extractionResults, nil
}

func (s *detectTextService) GetResultByID(id string) (*models.ExtractionResult, error) {
	result, err := s.extractionResultRepository.GetByID(s.db, id)
	if err != nil {
		return nil, err
	}
	return s.buildExtractionResultResponse(result), nil
}

func (s *detectTextService) DetectTexts(file *os.File) error {
	id := utils.NewULID()
	key := fmt.Sprintf("%s/%s", id, filepath.Base(file.Name()))

	err := s.storageAPIClient.UploadFile(key, file)
	if err != nil {
		return err
	}

	var result *entities.ExtractionResult

	status := enums.ExtractionResultStatus_Succeeded
	defer func() {
		if result != nil {
			result.Status = status
			_ = s.extractionResultRepository.Update(s.db, result)
		}
	}()

	return s.db.Transaction(func(tx *gorm.DB) error {
		extractionResult := &entities.ExtractionResult{
			ID:       id,
			Status:   enums.ExtractionResultStatus_Runing,
			ImageKey: key,
		}
		err = s.extractionResultRepository.Create(s.db, extractionResult)
		if err != nil {
			return err
		}
		extractionResult, err = s.extractionResultRepository.GetByID(s.db, id)
		if err != nil {
			return err
		}
		result = extractionResult
		outputKey, err := s.visionAPIClient.DetectText(key)
		if err != nil {
			status = enums.ExtractionResultStatus_Failed
			return err
		}

		queryFiles, err := s.storageAPIClient.QueryFiles(outputKey)
		if err != nil {
			status = enums.ExtractionResultStatus_Failed
			return err
		}

		if len(queryFiles) == 0 {
			status = enums.ExtractionResultStatus_Failed
			return errors.New("output does not exist")
		}

		outputFile, err := s.storageAPIClient.DownloadFile(queryFiles[0])
		if err != nil {
			status = enums.ExtractionResultStatus_Failed
			return err
		}

		result.OutputKey = &queryFiles[0]
		err = s.extractionResultRepository.Update(s.db, result)
		if err != nil {
			status = enums.ExtractionResultStatus_Failed
			return err
		}
		extractionResult, err = s.extractionResultRepository.GetByID(s.db, id)
		if err != nil {
			status = enums.ExtractionResultStatus_Failed
			return err
		}
		result = extractionResult

		bytes, err := io.ReadAll(outputFile)
		if err != nil {
			status = enums.ExtractionResultStatus_Failed
			return err
		}

		var detectResponse googleModels.DetectTextResponses
		if err = json.Unmarshal(bytes, &detectResponse); err != nil {
			status = enums.ExtractionResultStatus_Failed
			return err
		}
		extractedTexts := make([]entities.ExtractedText, 0)
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
						extractedText := entities.ExtractedText{
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
		return s.extractionResultRepository.Delete(tx, id)
	})
}

func (s *detectTextService) buildExtractionResultResponse(entity *entities.ExtractionResult) *models.ExtractionResult {
	imageUri := ""
	var err error
	if len(entity.ImageKey) != 0 {
		imageUri, err = s.storageAPIClient.SignedURL(entity.ImageKey)
		if err != nil {
			s.logger.Error(err.Error())
		}
	}
	outputUri := ""
	if entity.OutputKey != nil {
		outputUri, err = s.storageAPIClient.SignedURL(*entity.OutputKey)
		if err != nil {
			s.logger.Error(err.Error())
		}
	}
	extractedTexts := make([]models.ExtractedText, 0)

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
	return &models.ExtractionResult{
		ID:             entity.ID,
		Status:         entity.Status,
		ImageUri:       imageUri,
		OutputUri:      outputUri,
		CreatedAt:      entity.CreatedAt.Unix(),
		UpdatedAt:      entity.UpdatedAt.Unix(),
		ExtractedTexts: extractedTexts,
	}
}

type detectTextService struct {
	storageAPIClient           clients.StorageAPIClient
	visionAPIClient            clients.VisionAPIClient
	extractionResultRepository repositories.ExtractionResultRepository
	extractedTextRepository    repositories.ExtractedTextRepository
	db                         *gorm.DB
	logger                     *zap.Logger
}
