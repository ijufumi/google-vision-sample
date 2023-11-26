package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/google-vision-sample/internal/common/context"
	"github.com/ijufumi/google-vision-sample/internal/common/loggers"
	"github.com/ijufumi/google-vision-sample/internal/common/utils"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := utils.NewULID()
		l := loggers.NewLoggerWithTraceID(traceID)
		context.SetLoggerToGinContext(ctx, l)
		start := time.Now()
		path := ctx.Request.URL.Path
		if ctx.Request.URL.RawQuery != "" {
			path = path + "?" + ctx.Request.URL.RawQuery
		}
		method := ctx.Request.Method
		clientIP := ctx.ClientIP()

		l.Info("[START  ]",
			zap.String("timestamp", start.Format("2006/01/02 - 15:04:05")),
			zap.String("clientIP", clientIP),
			zap.String("method", method),
			zap.String("path", path),
		)

		ctx.Next()

		now := time.Now()
		status := ctx.Writer.Status()
		l.Info("[END  ]",
			zap.String("timestamp", now.Format("2006/01/02 - 15:04:05")),
			zap.Int64("latency", now.Sub(start).Milliseconds()),
			zap.Int("status", status),
			zap.String("clientIP", clientIP),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("size", ctx.Writer.Size()),
		)
	}
}

func outputRequestLog(l *zap.Logger, req http.Request) {
}
