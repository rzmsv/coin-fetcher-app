package logger

import "go.uber.org/zap"

func LoggerZap() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	return logger
}
