package clients

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/ijufumi/google-vision-sample/internal/common/configs"
	"github.com/ijufumi/google-vision-sample/internal/common/utils"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities/enums"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/google/models/services"
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
	UploadFile(ctx context.Context, key string, file *os.File, contentType enums.ContentType) error
	DownloadFile(ctx context.Context, key string) (*os.File, error)
	QueryFiles(ctx context.Context, key string) ([]string, error)
	DeleteFile(ctx context.Context, key string) error
	SignedURL(ctx context.Context, key string) (string, error)
	UpdateContentType(ctx context.Context, key, contentType string) error
	SetupCORSOnBucket() error
}

func NewStorageAPIClient(config *configs.Config) StorageAPIClient {
	return &storageAPIClient{
		config: config,
	}
}

type storageAPIClient struct {
	baseClient
	config    *configs.Config
	jwtConfig *jwt.Config
}

func (c *storageAPIClient) UploadFile(ctx context.Context, key string, file *os.File, contentType enums.ContentType) error {
	client, err := c.newClient()
	if err != nil {
		return errors.Wrap(err, "StorageAPIClient#UploadFile")
	}

	return c.Process(ctx, func(logger *zap.Logger) error {
		defer func() {
			err := client.Close()
			if err != nil {
				logger.Error("client closing error", zap.Error(err))
			}
		}()

		object := client.Bucket(c.config.Google.Storage.Bucket).Object(key)
		storageWriter := object.NewWriter(ctx)
		storageWriter.ContentType = string(contentType)
		defer func() {
			err := storageWriter.Close()
			if err != nil {
				logger.Error("writer closing error", zap.Error(err))
			}
		}()

		if size, err := io.Copy(storageWriter, file); err != nil {
			return errors.Wrap(err, "StorageAPIClient#UploadFile")
		} else if size == 0 {
			return errors.New("written size was 0")
		}

		return nil
	})
}

func (c *storageAPIClient) DownloadFile(ctx context.Context, key string) (*os.File, error) {
	client, err := c.newClient()
	if err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#DownloadFile")
	}
	defer func() {
		_ = client.Close()
	}()

	var tempFile *os.File
	err = c.Process(ctx, func(logger *zap.Logger) error {
		logger.Info(fmt.Sprintf("key is %s", key))
		object := client.Bucket(c.config.Google.Storage.Bucket).Object(key)
		storageReader, err := object.NewReader(ctx)
		if err != nil {
			return errors.Wrap(err, "StorageAPIClient#DownloadFile#NewReader")
		}
		defer func() {
			_ = storageReader.Close()
		}()
		tempFile, err = utils.NewTempFile()
		if err != nil {
			return errors.Wrap(err, "StorageAPIClient#DownloadFile#NewTempFile")
		}
		if _, err = io.Copy(tempFile, storageReader); err != nil {
			return errors.Wrap(err, "StorageAPIClient#DownloadFile#Copy")
		}
		_, err = tempFile.Seek(0, 0)
		if err != nil {
			return errors.Wrap(err, "StorageAPIClient#DownloadFile#Seek")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}

func (c *storageAPIClient) SignedURL(ctx context.Context, key string) (string, error) {
	client, err := c.newClient()
	if err != nil {
		return "", errors.Wrap(err, "StorageAPIClient#SignedURL")
	}
	signedURL := ""
	err = c.Process(ctx, func(logger *zap.Logger) error {
		defer func() {
			_ = client.Close()
		}()
		option := &storage.SignedURLOptions{
			Scheme:     storage.SigningSchemeV4,
			PrivateKey: c.jwtConfig.PrivateKey,
			Expires:    time.Now().UTC().Add(time.Duration(c.config.Google.Storage.SignedURL.ExpireSec) * time.Second),
			Method:     http.MethodGet,
		}
		_signedURL, err := client.Bucket(c.config.Google.Storage.Bucket).SignedURL(key, option)
		if err != nil {
			return errors.Wrap(err, "StorageAPIClient#SignedURL")
		}
		signedURL = _signedURL
		logger.Info(fmt.Sprintf("[%s]signed url is %s", key, signedURL))
		return nil
	})

	if err != nil {
		return "", err
	}
	return signedURL, nil
}

func (c *storageAPIClient) QueryFiles(ctx context.Context, key string) ([]string, error) {
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
	objects := client.Bucket(c.config.Google.Storage.Bucket).Objects(ctx, query)
	files := make([]string, 0)
	for {
		obj, err := objects.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			return nil, errors.Wrap(err, "StorageAPIClient#QueryFiles#Next")
		}
		files = append(files, obj.Name)
	}

	return files, nil
}

func (c *storageAPIClient) DeleteFile(ctx context.Context, key string) error {
	client, err := c.newClient()
	if err != nil {
		return errors.Wrap(err, "StorageAPIClient#DeleteFile")
	}
	defer func() {
		_ = client.Close()
	}()
	return c.Process(ctx, func(logger *zap.Logger) error {
		logger.Info(fmt.Sprintf("key is %s", key))
		object := client.Bucket(c.config.Google.Storage.Bucket).Object(key)
		err = object.Delete(ctx)
		if err != nil {
			return errors.Wrap(err, "StorageAPIClient#DeleteFile#Delete")
		}
		return nil
	})
}

func (c *storageAPIClient) UpdateContentType(ctx context.Context, key, contentType string) error {
	client, err := c.newClient()
	if err != nil {
		return errors.Wrap(err, "StorageAPIClient#UpdateContentType")
	}
	defer func() {
		_ = client.Close()
	}()
	return c.Process(ctx, func(logger *zap.Logger) error {
		logger.Info(fmt.Sprintf("key is %s", key))
		object := client.Bucket(c.config.Google.Storage.Bucket).Object(key)
		_, err = object.Update(ctx, storage.ObjectAttrsToUpdate{ContentType: contentType})
		if err != nil {
			return errors.Wrap(err, "StorageAPIClient#UpdateContentType#Update")
		}
		return nil
	})
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
	option := services.GetCredentialOption(c.config)
	service, err := storage.NewClient(context.Background(), option)
	if err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#newClient")
	}
	jwtConfig, err := services.GetJWTConfig(c.config)
	if err != nil {
		return nil, errors.Wrap(err, "StorageAPIClient#GetJWTConfig")
	}
	c.jwtConfig = jwtConfig
	return service, nil
}

func makeToGCSUri(bucket, key string) string {
	return fmt.Sprintf("gs://%s/%s", bucket, key)
}
