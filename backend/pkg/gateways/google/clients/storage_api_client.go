package clients

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/options"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
	"io"
	"os"
)

type StorageAPIClient interface {
	UploadFile(key string, file *os.File) error
}

func NewStorageAPIClient(config *configs.Config) StorageAPIClient {
	return &storageAPIClient{
		config: config,
	}
}

type storageAPIClient struct {
	config *configs.Config
}

func (c *storageAPIClient) UploadFile(key string, file *os.File) error {
	client := c.newClient()
	defer func() {
		_ = client.Close()
	}()

	object := client.Bucket(c.config.Google.Storage.Bucket).Object(key)
	storageWriter := object.NewWriter(context.Background())
	defer func() {
		_ = storageWriter.Close()
	}()

	if _, err := io.Copy(storageWriter, file); err != nil {
		return err
	}
	return nil
}

func (c *storageAPIClient) DownloadFile(key string) (*os.File, error) {
	client := c.newClient()
	defer func() {
		_ = client.Close()
	}()
	object := client.Bucket(c.config.Google.Storage.Bucket).Object(key)
	storageReader, err := object.NewReader(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = storageReader.Close()
	}()
	tempFile, err := utils.NewTempFile()
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(tempFile, storageReader); err != nil {
		return nil, err
	}
	_, err = tempFile.Seek(0, 0)
	if _, err = io.Copy(tempFile, storageReader); err != nil {
		return nil, err
	}
	return tempFile, nil
}

func (c *storageAPIClient) newClient() *storage.Client {
	option := options.GetCredentialOption(c.config)
	service, err := storage.NewClient(context.Background(), option)
	if err != nil {
		panic(err)
	}
	return service
}
