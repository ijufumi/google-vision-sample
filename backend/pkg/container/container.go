package container

import (
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/db"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/repositories"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/google/clients"
	"github.com/ijufumi/google-vision-sample/pkg/http/handlers"
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
	err := c.container.Provide(configs.New)
	if err != nil {
		fmt.Println(err)
	}
	// handlers
	err = c.container.Provide(handlers.NewHealthHandler)
	if err != nil {
		fmt.Println(err)
	}
	err = c.container.Provide(handlers.NewDetectTextHandler)
	if err != nil {
		fmt.Println(err)
	}
	err = c.container.Provide(handlers.NewSignedURL)
	if err != nil {
		fmt.Println(err)
	}
	err = c.container.Provide(handlers.NewConfigsHandler)
	if err != nil {
		fmt.Println(err)
	}
	// services
	err = c.container.Provide(services.NewDetectTextService)
	if err != nil {
		fmt.Println(err)
	}
	err = c.container.Provide(services.NewConfigurationService)
	if err != nil {
		fmt.Println(err)
	}
	err = c.container.Provide(services.NewImageConversionService)
	if err != nil {
		fmt.Println(err)
	}
	// database
	err = c.container.Provide(db.NewDB)
	if err != nil {
		fmt.Println(err)
	}
	err = c.container.Provide(repositories.NewJobRepository)
	if err != nil {
		fmt.Println(err)
	}
	err = c.container.Provide(repositories.NewExtractedTextRepository)
	if err != nil {
		fmt.Println(err)
	}
	err = c.container.Provide(repositories.NewInputFileRepository)
	if err != nil {
		fmt.Println(err)
	}
	err = c.container.Provide(repositories.NewOutputFileRepository)
	if err != nil {
		fmt.Println(err)
	}
	// google
	err = c.container.Provide(clients.NewStorageAPIClient)
	if err != nil {
		fmt.Println(err)
	}
	err = c.container.Provide(clients.NewVisionAPIClient)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *container) Container() *dig.Container {
	return c.container
}
