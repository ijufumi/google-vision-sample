package clients

import (
	vision "cloud.google.com/go/vision/v2/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"context"
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/options"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type VisionAPIClient interface {
	DetectText(key string) (string, error)
}

func NewVisionAPIClient(config *configs.Config, logger *zap.Logger) VisionAPIClient {
	return &visionAPIClient{
		config: config,
		logger: logger,
	}
}

func (c *visionAPIClient) DetectText(key string) (string, error) {
	client, err := c.newClient()
	if err != nil {
		return "", errors.Wrap(err, "VisionAPIClient#DetectText")
	}

	defer func() {
		_ = client.Close()
	}()

	imageUri := MakeToGCSUri(c.config.Google.Storage.Bucket, key)
	outputKey := fmt.Sprintf("%s-output-%d.json", key, time.Now().UTC().Unix())
	outputUri := fmt.Sprintf("gs://%s/%s", c.config.Google.Storage.Bucket, outputKey)

	c.logger.Info(fmt.Sprintf("imageUri is %s", imageUri))
	c.logger.Info(fmt.Sprintf("outputUri is %s", outputUri))
	request := &visionpb.AsyncBatchAnnotateImagesRequest{
		Requests: []*visionpb.AnnotateImageRequest{
			&visionpb.AnnotateImageRequest{
				Image: &visionpb.Image{
					Source: &visionpb.ImageSource{ImageUri: imageUri},
				},
				Features: []*visionpb.Feature{
					&visionpb.Feature{Type: visionpb.Feature_TEXT_DETECTION},
				},
			},
		},
		OutputConfig: &visionpb.OutputConfig{
			GcsDestination: &visionpb.GcsDestination{
				Uri: outputUri,
			},
		},
	}
	operation, err := client.AsyncBatchAnnotateImages(context.Background(), request)
	if err != nil {
		return "", errors.Wrap(err, "VisionAPIClient#DetectText")
	}
	response, err := operation.Wait(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "VisionAPIClient#DetectText")
	}

	c.logger.Info(fmt.Sprintf("%+v", response))

	return outputKey, nil
}

type visionAPIClient struct {
	config *configs.Config
	logger *zap.Logger
}

func (c *visionAPIClient) newClient() (*vision.ImageAnnotatorClient, error) {
	option := options.GetCredentialOption(c.config)
	service, err := vision.NewImageAnnotatorClient(context.Background(), option)
	if err != nil {
		return nil, errors.Wrap(err, "VisionAPIClient#newClient")
	}
	return service, nil
}
