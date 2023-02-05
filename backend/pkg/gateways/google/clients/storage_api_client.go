package clients

import (
	"context"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/options"
	"google.golang.org/api/storage/v1"
)

type StorageAPIClient interface {
}

func NewStorageAPIClient(config *configs.Config) StorageAPIClient {
	return &storageAPIClient{
		config: config,
	}
}

type storageAPIClient struct {
	config *configs.Config
}

func (c *storageAPIClient) newService() *storage.Service {
	option := options.GetCredentialOption(c.config)
	service, err := storage.NewService(context.Background(), option)
	if err != nil {
		panic(err)
	}
	return service
}
