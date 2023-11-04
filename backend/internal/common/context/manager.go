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

func GetDBFromGinContext(ctx *gin.Context) *gorm.DB {
	d, _ := ctx.Get("db")
	return d.(*gorm.DB)
}

func GetContext(ctx *gin.Context) context.Context {
	l := GetLoggerFromGinContext(ctx)
	d := GetDBFromGinContext(ctx)
	_ctx := context.WithValue(ctx.Request.Context(), "logger", l)
	return context.WithValue(_ctx, "db", d)
}

func GetLogger(ctx context.Context) *zap.Logger {
	l := ctx.Value("logger")
	return l.(*zap.Logger)
}

func GetDB(ctx context.Context) *gorm.DB {
	d := ctx.Value("db")
	return d.(*gorm.DB)
}
