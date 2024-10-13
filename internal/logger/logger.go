package logger

import (
	"go.uber.org/zap"
)

func NewLogger(level string) *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}
