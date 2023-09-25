package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/loggers"
)

func SetLoggerToGinContext(ctx *gin.Context, logger *loggers.Logger) {
	ctx.Set("logger", logger)
}

func GetLoggerFromGinContext(ctx *gin.Context) *loggers.Logger {
	l, _ := ctx.Get("logger")
	return l.(*loggers.Logger)
}

func GetContextWithLogger(ctx *gin.Context) context.Context {
	l := GetLoggerFromGinContext(ctx)
	return context.WithValue(ctx.Request.Context(), "logger", l)
}

func GetLogger(ctx context.Context) *loggers.Logger {
	l := ctx.Value("logger")
	return l.(*loggers.Logger)
}
