package services

import (
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/repositories"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/clients"
	"github.com/ijufumi/google-vision-sample/pkg/models"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
	"gorm.io/gorm"
	"os"
)

type DetectTextService interface {
	GetResults() ([]models.ExtractionResult, error)
	DetectTexts(file *os.File) error
}

func NewDetectTextService(
	storageAPIClient clients.StorageAPIClient,
	visionAPIClient clients.VisionAPIClient,
	extractionResultRepository repositories.ExtractionResultRepository,
	extractedTextRepository repositories.ExtractedTextRepository,
	db *gorm.DB,
) DetectTextService {
	return &detectTextService{
		storageAPIClient:           storageAPIClient,
		visionAPIClient:            visionAPIClient,
		extractionResultRepository: extractionResultRepository,
		extractedTextRepository:    extractedTextRepository,
		db:                         db,
	}
}

func (s *detectTextService) GetResults() ([]models.ExtractionResult, error) {
	results, err := s.extractionResultRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var extractionResults []models.ExtractionResult
	for _, result := range results {
		extractionResults = append(extractionResults, models.ExtractionResult{
			ID:        result.ID,
			Status:    result.Status,
			ImageUri:  result.ImageUri,
			OutputUri: result.OutputUri,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		})
	}
	return extractionResults, nil
}

func (s *detectTextService) DetectTexts(file *os.File) error {
	id := utils.NewULID()
	key := fmt.Sprintf("%s/%s", id, utils.NewRandomString(10))

	err := s.storageAPIClient.UploadFile(key, file)
	if err != nil {
		return nil
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		result := entities.ExtractionResult{
			ID:     id,
			Status: enums.ExtractionResultStatus_Runing,
		}
		err := s.extractionResultRepository.Create(result)
		if err != nil {
			return err
		}

		return nil
	})
}

type detectTextService struct {
	storageAPIClient           clients.StorageAPIClient
	visionAPIClient            clients.VisionAPIClient
	extractionResultRepository repositories.ExtractionResultRepository
	extractedTextRepository    repositories.ExtractedTextRepository
	db                         *gorm.DB
}
