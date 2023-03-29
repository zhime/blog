package core

import (
	"go.uber.org/zap"
)

func Zap() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}
