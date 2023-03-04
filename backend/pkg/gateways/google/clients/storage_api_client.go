package clients

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/options"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/iterator"
	"io"
	"net/http"
	"os"
	"time"
)

type StorageAPIClient interface {
	UploadFile(key string, file *os.File) error
	DownloadFile(key string) (*os.File, error)
	QueryFiles(key string) ([]string, error)
	DeleteFile(key string) error
	SignedURL(key string) (string, error)
	SetupCORSOnBucket() error
}

func NewStorageAPIClient(config *configs.Config, logger *zap.Logger) StorageAPIClient {
	return &storageAPIClient{
		config: config,
		logger: logger,
	}
}

type storageAPIClient struct {
	logger    *zap.Logger
	config    *configs.Config
	jwtConfig *jwt.Config
}

func (c *storageAPIClient) UploadFile(key string, file *os.File) error {
	client, err := c.newClient()
	if err != nil {
		return errors.Wrap(err, "StorageAPIClient#UploadFile")
	}
	defer func() {
		_ = client.Close()
	}()

	object := client.Bucket(c.config.Google.Storage.Bucket).Object(key)
	storageWriter := object.NewWriter(context.Background())
	defer func() {
		_ = storageWriter.Close()
	}()

	if _, err := io.Copy(storageWriter, file); err != nil {
		return errors.Wrap(err, "StorageAPIClient#UploadFile")
	}
	return nil
}

func (c *storageAPIClient) DownloadFile(key string) (*os.File, error) {
	client, err := c.newClient()
	if err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#DownloadFile")
	}
	defer func() {
		_ = client.Close()
	}()
	fmt.Println(fmt.Sprintf("key is %s", key))
	object := client.Bucket(c.config.Google.Storage.Bucket).Object(key)
	storageReader, err := object.NewReader(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#DownloadFile#NewReader")
	}
	defer func() {
		_ = storageReader.Close()
	}()
	tempFile, err := utils.NewTempFile()
	if err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#DownloadFile#NewTempFile")
	}
	if _, err = io.Copy(tempFile, storageReader); err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#DownloadFile#Copy")
	}
	_, err = tempFile.Seek(0, 0)
	if err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#DownloadFile#Seek")
	}
	return tempFile, nil
}

func (c *storageAPIClient) SignedURL(key string) (string, error) {
	client, err := c.newClient()
	if err != nil {
		return "", errors.Wrap(err, "StorageAPIClient#SignedURL")
	}
	defer func() {
		_ = client.Close()
	}()
	option := &storage.SignedURLOptions{
		Scheme:     storage.SigningSchemeV4,
		PrivateKey: c.jwtConfig.PrivateKey,
		Expires:    time.Now().UTC().Add(time.Duration(c.config.Google.Storage.SignedURL.ExpireSec) * time.Second),
		Method:     http.MethodGet,
	}
	signedURL, err := client.Bucket(c.config.Google.Storage.Bucket).SignedURL(key, option)
	if err != nil {
		return "", errors.Wrap(err, "StorageAPIClient#SignedURL")
	}
	c.logger.Info(fmt.Sprintf("[%s]signed url is %s", key, signedURL))
	return signedURL, nil
}

func (c *storageAPIClient) QueryFiles(key string) ([]string, error) {
	client, err := c.newClient()
	if err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#QueryFiles")
	}
	defer func() {
		_ = client.Close()
	}()
	query := &storage.Query{
		Prefix: key,
	}
	objects := client.Bucket(c.config.Google.Storage.Bucket).Objects(context.Background(), query)
	files := make([]string, 0)
	for {
		obj, err := objects.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, errors.Wrap(err, "StorageAPIClient#QueryFiles#Next")
		}
		files = append(files, obj.Name)
	}

	return files, nil
}

func (c *storageAPIClient) DeleteFile(key string) error {
	client, err := c.newClient()
	if err != nil {
		return errors.Wrap(err, "StorageAPIClient#DeleteFile")
	}
	defer func() {
		_ = client.Close()
	}()
	object := client.Bucket(c.config.Google.Storage.Bucket).Object(key)
	err = object.Delete(context.Background())
	if err != nil {
		return errors.Wrap(err, "StorageAPIClient#DeleteFile#Delete")
	}
	return nil
}

func (c *storageAPIClient) SetupCORSOnBucket() error {
	client, err := c.newClient()
	if err != nil {
		return errors.Wrap(err, "StorageAPIClient#DeleteFile")
	}
	defer func() {
		_ = client.Close()
	}()
	bucket := client.Bucket(c.config.Google.Storage.Bucket)
	corsConfig := storage.BucketAttrsToUpdate{
		CORS: []storage.CORS{
			{
				MaxAge:          time.Duration(c.config.Google.Storage.SignedURL.ExpireSec) * time.Second,
				Methods:         []string{"GET"},
				Origins:         []string{"*"},
				ResponseHeaders: []string{"Access-Control-Allow-Origin"},
			},
		},
	}
	_, err = bucket.Update(context.Background(), corsConfig)
	return err
}

func (c *storageAPIClient) newClient() (*storage.Client, error) {
	option := options.GetCredentialOption(c.config)
	service, err := storage.NewClient(context.Background(), option)
	if err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#newClient")
	}
	jwtConfig, err := options.GetJWTConfig(c.config)
	if err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#GetJWTConfig")
	}
	c.jwtConfig = jwtConfig
	return service, nil
}

func MakeToGCSUri(bucket, key string) string {
	return fmt.Sprintf("gs://%s/%s", bucket, key)
}
