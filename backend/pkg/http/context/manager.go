package context

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/pkg/loggers"
)

func SetLogger(ctx *gin.Context, logger *loggers.Logger) {
	ctx.Set("logger", logger)
}

func GetLogger(ctx *gin.Context) *loggers.Logger {
	l, _ := ctx.Get("logger")
	return l.(*loggers.Logger)
}
