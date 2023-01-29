package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/container"
)

type Router interface {
	Run() error
}

func NewRouter(container container.Container) Router {
	r := gin.Default()
	return &router{engine: r}
}

type router struct {
	engine *gin.Engine
}

func (r *router) Run() error {
	return r.engine.Run()
}
