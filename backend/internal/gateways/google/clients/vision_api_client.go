package clients

import (
	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"context"
	"fmt"
	"github.com/ijufumi/google-vision-sample/internal/common/configs"
	"github.com/ijufumi/google-vision-sample/internal/gateways/google/options"
	"github.com/pkg/errors"
	"time"
)

type VisionAPIClient interface {
	DetectText(key string) (string, error)
}

func NewVisionAPIClient(config *configs.Config) VisionAPIClient {
	return &visionAPIClient{
		config: config,
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
	outputUri := MakeToGCSUri(c.config.Google.Storage.Bucket, outputKey)

	// c.logger.Info(fmt.Sprintf("imageUri is %s", imageUri))
	// c.logger.Info(fmt.Sprintf("outputUri is %s", outputUri))
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
		return "", errors.Wrap(err, "VisionAPIClient#DetectText")
	}
	response, err := operation.Wait(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "VisionAPIClient#DetectText")
	}

	fmt.Println(response)
	//c.logger.Info(fmt.Sprintf("%+v", response))

	return outputKey, nil
}

type visionAPIClient struct {
	config *configs.Config
}

func (c *visionAPIClient) newClient() (*vision.ImageAnnotatorClient, error) {
	option := options.GetCredentialOption(c.config)
	service, err := vision.NewImageAnnotatorClient(context.Background(), option)
	if err != nil {
		return nil, errors.Wrap(err, "VisionAPIClient#newClient")
	}
	return service, nil
}
