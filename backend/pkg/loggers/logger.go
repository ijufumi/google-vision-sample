package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *Logger {
	return &Logger{buildZapLogger()}
}

func buildZapLogger() zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	l, _ := loggerConfig.Build()
	return *l
}

type Logger struct {
	zap.Logger
}
