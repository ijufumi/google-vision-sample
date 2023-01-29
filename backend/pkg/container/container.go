package container

import "go.uber.org/dig"

type Container interface {
}

func NewConainer() Container {
	return &container{}
}

type container struct {
	container *dig.Container
}
