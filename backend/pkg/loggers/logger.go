package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger = buildZapLogger()

func NewLogger() *zap.Logger {
	return zapLogger
}

func NewLoggerWithTraceID(traceID string) *zap.Logger {
	return zapLogger.With(zap.String("trace_id", traceID))
}

func buildZapLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	l, _ := loggerConfig.Build()
	return l
}
