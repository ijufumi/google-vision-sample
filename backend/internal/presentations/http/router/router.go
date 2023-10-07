package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/internal/common/container"
	handlers2 "github.com/ijufumi/google-vision-sample/internal/presentations/http/handlers"
	middlewares2 "github.com/ijufumi/google-vision-sample/internal/presentations/http/middlewares"
)

type Router interface {
	Run() error
}

func NewRouter(c container.Container) Router {
	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false
	r.Use(middlewares2.ResponseHeaders())
	r.Use(middlewares2.CORS())
	r.Use(middlewares2.Logger())

	api := r.Group("api")
	{
		healthHandler := container.Invoke[handlers2.HealthHandler](c)
		api.GET("/health", healthHandler.Get)

		v1 := api.Group("v1")
		detectTextHandler := container.Invoke[handlers2.DetectTextHandler](c)
		v1.GET("/detect_texts", detectTextHandler.Gets)
		v1.GET("/detect_texts/:id", detectTextHandler.GetByID)
		v1.POST("/detect_texts", detectTextHandler.Create)
		v1.DELETE("/detect_texts/:id", detectTextHandler.Delete)

		signedURLHandler := container.Invoke[handlers2.SignedURLHandler](c)
		v1.GET("/signed_urls", signedURLHandler.GetByKey)

		configsHandler := container.Invoke[handlers2.ConfigsHandler](c)
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
