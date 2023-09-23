package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/loggers"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		l := loggers.NewLogger()
		ctx.Set("logger", l)
	}
}
