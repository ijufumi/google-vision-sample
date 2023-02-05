package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/container"
	"github.com/ijufumi/google-vision-sample/pkg/http/handlers"
)

type Router interface {
	Run() error
}

func NewRouter(c container.Container) Router {
	r := gin.Default()

	api := r.Group("api/v1")
	{
		healthHandler := container.Invoke[handlers.HealthHandler](c)
		api.GET("/health", healthHandler.Get)
		detectTextHandler := container.Invoke[handlers.DetectTextHandler](c)
		api.GET("/detect_texts", detectTextHandler.Get)
		api.POST("/detect_texts", detectTextHandler.Post)
	}
	return &router{engine: r}
}

type router struct {
	engine *gin.Engine
}

func (r *router) Run() error {
	return r.engine.Run()
}
