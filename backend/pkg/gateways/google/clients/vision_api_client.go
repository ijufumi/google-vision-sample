package clients

import (
	vision "cloud.google.com/go/vision/v2/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"context"
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/options"
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
		return "", err
	}

	defer func() {
		_ = client.Close()
	}()

	imageUri := fmt.Sprintf("gs://%s/%s", c.config.Google.Storage.Bucket, key)
	outputUri := fmt.Sprintf("gs://%s/%s-output-%d.json", c.config.Google.Storage.Bucket, key, time.Now().UTC().Unix())
	request := &visionpb.AsyncBatchAnnotateImagesRequest{
		Requests: []*visionpb.AnnotateImageRequest{
			&visionpb.AnnotateImageRequest{
				Image: &visionpb.Image{
					Source: &visionpb.ImageSource{GcsImageUri: imageUri},
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
		return "", err
	}
	_, err = operation.Wait(context.Background())
	if err != nil {
		return "", err
	}

	return outputUri, nil
}

type visionAPIClient struct {
	config *configs.Config
}

func (c *visionAPIClient) newClient() (*vision.ImageAnnotatorClient, error) {
	option, err := options.GetCredentialOption(c.config)
	if err != nil {
		return nil, err
	}
	service, err := vision.NewImageAnnotatorClient(context.Background(), option)
	if err != nil {
		return nil, err
	}
	return service, nil
}
