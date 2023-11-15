package clients

import (
	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"context"
	"fmt"
	"github.com/ijufumi/google-vision-sample/internal/common/configs"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/google/models/services"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type VisionAPIClient interface {
	DetectText(ctx context.Context, key string) (string, error)
}

func NewVisionAPIClient(config *configs.Config) VisionAPIClient {
	return &visionAPIClient{
		config: config,
	}
}

func (c *visionAPIClient) DetectText(ctx context.Context, key string) (string, error) {
	client, err := c.newClient()
	if err != nil {
		return "", errors.Wrap(err, "VisionAPIClient#DetectText")
	}

	defer func() {
		_ = client.Close()
	}()

	imageUri := makeToGCSUri(c.config.Google.Storage.Bucket, key)
	outputKey := fmt.Sprintf("%s-output-%d.json", key, time.Now().UTC().Unix())
	outputUri := makeToGCSUri(c.config.Google.Storage.Bucket, outputKey)
	err = c.Process(ctx, func(logger *zap.Logger) error {

		logger.Info(fmt.Sprintf("imageUri is %s", imageUri))
		logger.Info(fmt.Sprintf("outputUri is %s", outputUri))
		request := &visionpb.AsyncBatchAnnotateImagesRequest{
			Requests: []*visionpb.AnnotateImageRequest{
				&visionpb.AnnotateImageRequest{
					Image: &visionpb.Image{
						Source: &visionpb.ImageSource{ImageUri: imageUri},
					},
					Features: []*visionpb.Feature{
						&visionpb.Feature{Type: visionpb.Feature_TEXT_DETECTION},
					},
					ImageContext: &visionpb.ImageContext{
						LanguageHints: []string{"ja"},
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
			return errors.Wrap(err, "VisionAPIClient#DetectText")
		}
		response, err := operation.Wait(context.Background())
		if err != nil {
			return errors.Wrap(err, "VisionAPIClient#DetectText")
		}

		if logger.Level() == zap.DebugLevel {
			logger.Debug(fmt.Sprintf("%+v", response))
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return outputKey, nil
}

type visionAPIClient struct {
	baseClient
	config *configs.Config
}

func (c *visionAPIClient) newClient() (*vision.ImageAnnotatorClient, error) {
	option := services.GetCredentialOption(c.config)
	service, err := vision.NewImageAnnotatorClient(context.Background(), option)
	if err != nil {
		return nil, errors.Wrap(err, "VisionAPIClient#newClient")
	}
	return service, nil
}
