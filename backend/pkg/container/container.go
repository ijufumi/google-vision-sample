package container

import (
	"github.com/ijufumi/google-vision-sample/pkg/http/handlers"
	"go.uber.org/dig"
)

type Container interface {
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
