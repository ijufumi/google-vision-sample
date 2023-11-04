package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetDBToGinContext(ctx *gin.Context, db *gorm.DB) {
	ctx.Set("db", db)
}

func SetLoggerToGinContext(ctx *gin.Context, logger *zap.Logger) {
	ctx.Set("logger", logger)
}

func GetLoggerFromGinContext(ctx *gin.Context) *zap.Logger {
	l, _ := ctx.Get("logger")
	return l.(*zap.Logger)
}

func GetContextWithLogger(ctx *gin.Context) context.Context {
	l := GetLoggerFromGinContext(ctx)
	return context.WithValue(ctx.Request.Context(), "logger", l)
}

func GetLogger(ctx context.Context) *zap.Logger {
	l := ctx.Value("logger")
	return l.(*zap.Logger)
}
