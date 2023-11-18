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

func SetLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, "logger", logger)
}

func GetDB(ctx context.Context) *gorm.DB {
	d := ctx.Value("tx")
	if d != nil {
		return d.(*gorm.DB)
	}
	d = ctx.Value("db")
	return d.(*gorm.DB)
}

func SetDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, "db", db)
}

func SetTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, "tx", tx)
}

func NewContext(ctx context.Context) context.Context {
	ctx2 := context.Background()
	tx := ctx.Value("tx")
	if tx != nil {
		ctx2 = SetTx(ctx2, tx.(*gorm.DB))
	}
	db := ctx.Value("db")
	if db != nil {
		ctx2 = SetDB(ctx2, db.(*gorm.DB))
	}
	l := ctx.Value("logger")
	if l != nil {
		ctx2 = SetLogger(ctx2, l.(*zap.Logger))
	}
	return ctx2
}
