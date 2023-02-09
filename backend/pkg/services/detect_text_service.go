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

	var detectResponse googleModels.DetectTextResponses
	if err = json.Unmarshal(bytes, &detectResponse); err != nil {
		status = enums.ExtractionResultStatus_Failed
		return err
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
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
						top := utils.MaxInArray(vertices[0].X)
						bottom := utils.MaxInArray(vertices[0].X)
						left := utils.MaxInArray(vertices[0].X)
						right := utils.MaxInArray(vertices[0].X)
						extractedText := entities.ExtractedText{
							ID:                 utils.NewULID(),
							ExtractionResultID: result.ID,
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
		err := s.extractedTextRepository.Create(tx, extractedTexts...)
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
