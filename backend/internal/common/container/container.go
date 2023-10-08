package container

import (
	"github.com/ijufumi/google-vision-sample/internal/common/configs"
	"github.com/ijufumi/google-vision-sample/internal/common/loggers"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/db"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/repositories"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/google/clients"
	"github.com/ijufumi/google-vision-sample/internal/models/service"
	"github.com/ijufumi/google-vision-sample/internal/presentations/handlers"
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
		logger.Panic(err.Error())
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
	_ = c.container.Provide(handlers.NewHealthHandler)
	_ = c.container.Provide(handlers.NewDetectTextHandler)
	_ = c.container.Provide(handlers.NewSignedURL)
	_ = c.container.Provide(handlers.NewConfigsHandler)
	// services
	_ = c.container.Provide(service.NewDetectTextService)
	_ = c.container.Provide(service.NewConfigurationService)
	_ = c.container.Provide(service.NewImageConversionService)
	// database
	_ = c.container.Provide(db.NewDB)
	_ = c.container.Provide(repositories.NewJobRepository)
	_ = c.container.Provide(repositories.NewExtractedTextRepository)
	_ = c.container.Provide(repositories.NewInputFileRepository)
	_ = c.container.Provide(repositories.NewOutputFileRepository)
	// google
	_ = c.container.Provide(clients.NewStorageAPIClient)
	_ = c.container.Provide(clients.NewVisionAPIClient)
}

func (c *container) Container() *dig.Container {
	return c.container
}
