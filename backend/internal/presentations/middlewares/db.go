package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/internal/common/context"
	"github.com/ijufumi/google-vision-sample/internal/common/loggers"
	"github.com/ijufumi/google-vision-sample/internal/common/utils"
)

func DB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := utils.NewULID()
		l := loggers.NewLoggerWithTraceID(traceID)
		context.SetLoggerToGinContext(ctx, l)
	}
}
