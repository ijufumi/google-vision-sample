package services

import (
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities/enums"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/repositories"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/clients"
	"github.com/ijufumi/google-vision-sample/pkg/models"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
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
	key := fmt.Sprintf("%s/%s", id, filepath.Base(file.Name()))

	err := s.storageAPIClient.UploadFile(key, file)
	if err != nil {
		return err
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

	queryFiles, err := s.storageAPIClient.QueryFiles(outputKey)
	if err != nil {
		status = enums.ExtractionResultStatus_Failed
		return err
	}

	if len(queryFiles) == 0 {
		return errors.New("output does not exist")
	}

	outputFile, err := s.storageAPIClient.DownloadFile(queryFiles[0])
	if err != nil {
		status = enums.ExtractionResultStatus_Failed
		return err
	}

	bytes, err := io.ReadAll(outputFile)
	if err != nil {
		status = enums.ExtractionResultStatus_Failed
		return err
	}

	fmt.Println(fmt.Sprintf("%+v", string(bytes)))

	return nil
}

type detectTextService struct {
	storageAPIClient           clients.StorageAPIClient
	visionAPIClient            clients.VisionAPIClient
	extractionResultRepository repositories.ExtractionResultRepository
	extractedTextRepository    repositories.ExtractedTextRepository
	db                         *gorm.DB
}
