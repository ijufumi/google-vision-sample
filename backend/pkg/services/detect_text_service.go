package services

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/clients"
	"github.com/ijufumi/google-vision-sample/pkg/models"
	"os"
)

type DetectTextService interface {
	GetResults() ([]models.ExtractionResult, error)
	DetectTexts(file *os.File) error
}

func NewDetectTextService(
	storageAPIClient clients.StorageAPIClient,
	visionAPIClient clients.VisionAPIClient,
) DetectTextService {
	return &detectTextService{
		storageAPIClient: storageAPIClient,
		visionAPIClient:  visionAPIClient,
	}
}

func (s *detectTextService) GetResults() ([]models.ExtractionResult, error) {
	return nil, nil
}

func (s *detectTextService) DetectTexts(file *os.File) error {
	return nil
}

type detectTextService struct {
	storageAPIClient clients.StorageAPIClient
	visionAPIClient  clients.VisionAPIClient
}
