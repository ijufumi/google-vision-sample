package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const contextKey_DB = "db"
const contextKey_Tx = "tx"
const contextKey_Logger = "logger"

func SetDBToGinContext(ctx *gin.Context, db *gorm.DB) {
	ctx.Set(contextKey_DB, db)
}

func SetLoggerToGinContext(ctx *gin.Context, logger *zap.Logger) {
	ctx.Set(contextKey_Logger, logger)
}

func GetLoggerFromGinContext(ctx *gin.Context) *zap.Logger {
	l, _ := ctx.Get(contextKey_Logger)
	return l.(*zap.Logger)
}

func GetDBFromGinContext(ctx *gin.Context) *gorm.DB {
	d, _ := ctx.Get(contextKey_DB)
	return d.(*gorm.DB)
}

func GetContext(ctx *gin.Context) context.Context {
	l := GetLoggerFromGinContext(ctx)
	d := GetDBFromGinContext(ctx)
	_ctx := context.WithValue(ctx.Request.Context(), contextKey_Logger, l)
	return context.WithValue(_ctx, contextKey_DB, d)
}

func GetLogger(ctx context.Context) *zap.Logger {
	return getValue[*zap.Logger](ctx, contextKey_Logger)
}

func GetDB(ctx context.Context) *gorm.DB {
	return getValue[*gorm.DB](ctx, contextKey_DB)
}

func GetTx(ctx context.Context) *gorm.DB {
	return getValue[*gorm.DB](ctx, contextKey_Tx)
}

func getValue[T any](ctx context.Context, contextKey string) T {
	v := ctx.Value(contextKey)
	return v.(T)
}

func SetLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, contextKey_Logger, logger)
}

func SetDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, contextKey_DB, db)
}

func SetTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, contextKey_Tx, tx)
}

func HasTx(ctx context.Context) bool {
	d := ctx.Value(contextKey_Tx)
	return d != nil
}

func NewContext(ctx context.Context) context.Context {
	ctx2 := context.Background()
	tx := ctx.Value(contextKey_Tx)
	if tx != nil {
		ctx2 = SetTx(ctx2, tx.(*gorm.DB))
	}
	db := ctx.Value(contextKey_DB)
	if db != nil {
		ctx2 = SetDB(ctx2, db.(*gorm.DB))
	}
	l := ctx.Value(contextKey_Logger)
	if l != nil {
		ctx2 = SetLogger(ctx2, l.(*zap.Logger))
	}
	return ctx2
}
