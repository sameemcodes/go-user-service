package durable

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{logger.Sugar()}
}
