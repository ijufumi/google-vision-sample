package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/internal/common/configs"
	"github.com/ijufumi/google-vision-sample/internal/common/container"
	"github.com/ijufumi/google-vision-sample/internal/presentations/handlers"
	"github.com/ijufumi/google-vision-sample/internal/presentations/middlewares"
)

type Router interface {
	Run() error
}

func NewRouter(c container.Container) Router {
	config := container.Invoke[*configs.Config](c)

	r := gin.New()
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false
	r.Use(gin.Recovery())
	r.Use(middlewares.ResponseHeaders())
	r.Use(middlewares.CORS())
	r.Use(middlewares.Logger())
	r.Use(middlewares.DB(config))

	api := r.Group("api")
	{
		healthHandler := container.Invoke[handlers.HealthHandler](c)
		api.GET("/health", healthHandler.Get)

		v1 := api.Group("v1")
		detectTextHandler := container.Invoke[handlers.DetectTextHandler](c)
		v1.GET("/detect_texts", detectTextHandler.Gets)
		v1.GET("/detect_texts/:id", detectTextHandler.GetByID)
		v1.POST("/detect_texts", detectTextHandler.Create)
		v1.DELETE("/detect_texts/:id", detectTextHandler.Delete)

		signedURLHandler := container.Invoke[handlers.SignedURLHandler](c)
		v1.GET("/signed_urls", signedURLHandler.GetByKey)

		configsHandler := container.Invoke[handlers.ConfigsHandler](c)
		v1.POST("/configs/cors", configsHandler.SetupCORS)
	}
	return &router{engine: r}
}

type router struct {
	engine *gin.Engine
}

func (r *router) Run() error {
	return r.engine.Run()
}
