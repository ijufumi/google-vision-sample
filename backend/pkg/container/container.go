package container

import (
	"github.com/ijufumi/google-vision-sample/pkg/http/handlers"
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

func NewConainer() Container {
	c := container{container: dig.New()}
	c.provide()
	return &c
}

type container struct {
	container *dig.Container
}

func (c *container) provide() {
	_ = c.container.Provide(handlers.NewHealthHandler())
}

func (c *container) Container() *dig.Container {
	return c.container
}
