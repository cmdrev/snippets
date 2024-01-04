package logger

import (
	prettyconsole "github.com/thessem/zap-prettyconsole"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogEnvironment string

const (
	LogEnvironmentProduction  LogEnvironment = "prod"
	LogEnvironmentDevelopment LogEnvironment = "dev"
)

type ZapLogConfig interface {
	Environment() LogEnvironment
	LogLevel() zapcore.Level
}

func NewZapLogger(config ZapLogConfig) (logger *zap.Logger, err error) {
	switch config.Environment() {
	case LogEnvironmentDevelopment:
		logger = prettyconsole.NewLogger(config.LogLevel())
	case LogEnvironmentProduction:
	default:
		logConfig := zap.NewProductionConfig()
		logConfig.Level = zap.NewAtomicLevelAt(config.LogLevel())
		logger, err = logConfig.Build()
	}
	return logger, err
}

func NewZapPrettyConsoleTextLogger() *zap.Logger {
	logger := prettyconsole.NewLogger(zap.DebugLevel)
	return logger
}
