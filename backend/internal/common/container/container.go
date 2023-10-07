package container

import (
	"github.com/ijufumi/google-vision-sample/internal/common/configs"
	"github.com/ijufumi/google-vision-sample/internal/common/loggers"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/db"
	repositories2 "github.com/ijufumi/google-vision-sample/internal/infrastructures/database/repositories"
	clients2 "github.com/ijufumi/google-vision-sample/internal/infrastructures/google/clients"
	"github.com/ijufumi/google-vision-sample/internal/models/service"
	handlers2 "github.com/ijufumi/google-vision-sample/internal/presentations/http/handlers"
	"go.uber.org/dig"
)

type Container interface {
	Container() *dig.Container
}

func Invoke[T any](container Container) T {
	logger := loggers.NewLogger()
	var result T
	err := container.Container().Invoke(func(_result T) {
		result = _result
	})
	if err != nil {
		logger.Error(err.Error())
	}

	return result
}

func NewContainer() Container {
	c := container{container: dig.New()}
	c.provide()
	return &c
}

type container struct {
	container *dig.Container
}

func (c *container) provide() {
	// logger
	_ = c.container.Provide(loggers.NewLogger)
	// config
	_ = c.container.Provide(configs.New)
	// handlers
	_ = c.container.Provide(handlers2.NewHealthHandler)
	_ = c.container.Provide(handlers2.NewDetectTextHandler)
	_ = c.container.Provide(handlers2.NewSignedURL)
	_ = c.container.Provide(handlers2.NewConfigsHandler)
	// services
	_ = c.container.Provide(service.NewDetectTextService)
	_ = c.container.Provide(service.NewConfigurationService)
	_ = c.container.Provide(service.NewImageConversionService)
	// database
	_ = c.container.Provide(db.NewDB)
	_ = c.container.Provide(repositories2.NewJobRepository)
	_ = c.container.Provide(repositories2.NewExtractedTextRepository)
	_ = c.container.Provide(repositories2.NewInputFileRepository)
	_ = c.container.Provide(repositories2.NewOutputFileRepository)
	// google
	_ = c.container.Provide(clients2.NewStorageAPIClient)
	_ = c.container.Provide(clients2.NewVisionAPIClient)
}

func (c *container) Container() *dig.Container {
	return c.container
}
