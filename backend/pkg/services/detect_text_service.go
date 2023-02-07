package services

import (
	"context"
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/repositories"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/clients"
	"github.com/ijufumi/google-vision-sample/pkg/models"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
	"google.golang.org/appengine/log"
	"gorm.io/gorm"
	"io"
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
	results, err := s.extractionResultRepository.GetAll(s.db)
	if err != nil {
		return nil, err
	}

	extractionResults := make([]models.ExtractionResult, 0)
	for _, result := range results {
		imageUri := ""
		if len(result.ImageKey) != 0 {
			imageUri, _ = s.storageAPIClient.SignedURL(result.ImageKey)
		}
		outputUri := ""
		if result.OutputKey != nil {
			outputUri, _ = s.storageAPIClient.SignedURL(*result.OutputKey)
		}
		extractionResults = append(extractionResults, models.ExtractionResult{
			ID:        result.ID,
			Status:    result.Status,
			ImageUri:  imageUri,
			OutputUri: outputUri,
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
	result := entities.ExtractionResult{
		ID:       id,
		Status:   enums.ExtractionResultStatus_Runing,
		ImageKey: key,
	}
	err = s.extractionResultRepository.Create(s.db, result)
	if err != nil {
		return err
	}
	status := enums.ExtractionResultStatus_Succeeded
	defer func() {
		result.Status = status
		_ = s.extractionResultRepository.Update(s.db, result)
	}()

	outputKey, err := s.visionAPIClient.DetectText(key)
	if err != nil {
		status = enums.ExtractionResultStatus_Failed
		return err
	}

	outputFile, err := s.storageAPIClient.DownloadFile(outputKey)
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(outputFile)
	if err != nil {
		return err
	}

	log.Infof(context.Background(), "%+v", string(bytes))

	return s.db.Transaction(func(tx *gorm.DB) error {
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
