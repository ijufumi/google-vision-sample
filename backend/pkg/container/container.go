package container

import (
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/db"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/repositories"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/clients"
	"github.com/ijufumi/google-vision-sample/pkg/http/handlers"
	"github.com/ijufumi/google-vision-sample/pkg/loggers"
	"github.com/ijufumi/google-vision-sample/pkg/services"
	"go.uber.org/dig"
)

type Container interface {
	Container() *dig.Container
}

func Invoke[T any](container Container) T {
	var result T
	_ = container.Container().Invoke(func(_result T) {
		result = _result
	})

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
	// config
	_ = c.container.Provide(configs.New)
	// logger
	_ = c.container.Provide(loggers.NewLogger)
	// handlers
	_ = c.container.Provide(handlers.NewHealthHandler)
	_ = c.container.Provide(handlers.NewDetectTextHandler)
	// services
	_ = c.container.Provide(services.NewDetectTextService)
	// database
	_ = c.container.Provide(db.NewDB)
	_ = c.container.Provide(repositories.NewExtractionResultRepository)
	_ = c.container.Provide(repositories.NewExtractedTextRepository)
	// google
	_ = c.container.Provide(clients.NewStorageAPIClient)
	_ = c.container.Provide(clients.NewVisionAPIClient)
}

func (c *container) Container() *dig.Container {
	return c.container
}
