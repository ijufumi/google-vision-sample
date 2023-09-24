package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger = buildZapLogger()

func NewLogger() *Logger {
	return &Logger{zapLogger}
}

func NewLoggerWithTraceID(traceID string) *Logger {
	return &Logger{zapLogger.With(zap.String("trace_id", traceID))}
}

func buildZapLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	l, _ := loggerConfig.Build()
	return l
}

type Logger struct {
	*zap.Logger
}
