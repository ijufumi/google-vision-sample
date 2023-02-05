package clients

import (
	"context"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/options"
	"google.golang.org/api/vision/v1"
)

type VisionAPIClient interface {
}

func NewVisionAPIClient(config *configs.Config) VisionAPIClient {
	return &visionAPIClient{
		config: config,
	}
}

type visionAPIClient struct {
	config *configs.Config
}

func (c *visionAPIClient) newService() *vision.Service {
	option := options.GetCredentialOption(c.config)
	service, err := vision.NewService(context.Background(), option)
	if err != nil {
		panic(err)
	}
	return service
}
