package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/container"
	"github.com/ijufumi/google-vision-sample/pkg/http/handlers"
)

type Router interface {
	Run() error
}

func NewRouter(c container.Container) Router {
	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	api := r.Group("api/v1")
	{
		healthHandler := container.Invoke[handlers.HealthHandler](c)
		api.GET("/health", healthHandler.Get)

		detectTextHandler := container.Invoke[handlers.DetectTextHandler](c)
		api.GET("/detect_texts", detectTextHandler.Gets)
		api.GET("/detect_texts/:id", detectTextHandler.GetByID)
		api.POST("/detect_texts", detectTextHandler.Create)
		api.DELETE("/detect_texts/:id", detectTextHandler.Delete)

		signedURLHandler := container.Invoke[handlers.SignedURLHandler](c)
		api.GET("/signed_urls", signedURLHandler.GetByKey)

		configsHandler := container.Invoke[handlers.ConfigsHandler](c)
		api.POST("/configs/cors", configsHandler.SetupCORS)
	}
	return &router{engine: r}
}

type router struct {
	engine *gin.Engine
}

func (r *router) Run() error {
	return r.engine.Run()
}
