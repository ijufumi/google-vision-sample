package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/http/context"
	"github.com/ijufumi/google-vision-sample/pkg/loggers"
	"github.com/ijufumi/google-vision-sample/pkg/utils"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := utils.NewULID()
		l := loggers.NewLoggerWithTraceID(traceID)
		context.SetLogger(ctx, l)
	}
}
