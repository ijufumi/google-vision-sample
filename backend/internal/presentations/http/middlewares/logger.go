package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/internal/common/loggers"
	"github.com/ijufumi/google-vision-sample/internal/common/utils"
	"github.com/ijufumi/google-vision-sample/internal/presentations/http/context"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := utils.NewULID()
		l := loggers.NewLoggerWithTraceID(traceID)
		context.SetLoggerToGinContext(ctx, l)
	}
}
