package service

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/google/clients"
)

type ConfigurationService interface {
	SetupCORS() error
}

type configurationService struct {
	storageAPIClient clients.StorageAPIClient
}

func NewConfigurationService(storageAPIClient clients.StorageAPIClient) ConfigurationService {
	return &configurationService{storageAPIClient: storageAPIClient}
}

func (s *configurationService) SetupCORS() error {
	return s.storageAPIClient.SetupCORSOnBucket()
}
