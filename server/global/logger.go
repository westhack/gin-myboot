package global

import "go.uber.org/zap"

func Debug(msg string, value interface{}) {
	Logger.Debug("debug", zap.Any(msg, value))
}

func Error(msg string, value interface{}) {
	Logger.Error("error", zap.Any(msg, value))
}

func Info(msg string, value interface{}) {
	Logger.Info("info", zap.Any(msg, value))
}
