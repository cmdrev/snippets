package config

import (
	"errors"
	"fmt"
	"github.com/cmdrev/snippets/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type ZapLogConfig struct {
	Environment string `yaml:"environment"`
	Level       string `yaml:"level"`
}

type ZapLog struct {
	logger.ZapLogConfig
	environment logger.LogEnvironment
	logLevel    zapcore.Level
}

func NewZapLogConfig(viperClient *viper.Viper) (logger.ZapLogConfig, error) {
	zapConfig := ZapLog{}
	configData := ZapLogConfig{}
	err := viperClient.UnmarshalKey("logging", &configData)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal zap log configuration: %w", err)
	}
	err = zapConfig.parseEnvironment(configData.Environment)
	if err != nil {
		return nil, fmt.Errorf("could not parse zap log configuration: %w", err)
	}
	err = zapConfig.parseLevel(configData.Level)
	if err != nil {
		return nil, fmt.Errorf("could not parse zap log configuration: %w", err)
	}
	return &zapConfig, nil
}

func (c *ZapLog) Environment() logger.LogEnvironment {
	return c.environment
}

func (c *ZapLog) LogLevel() zapcore.Level {
	return c.logLevel
}

func (c *ZapLog) parseEnvironment(environment string) error {
	switch strings.ToLower(environment) {
	case "dev":
		fallthrough
	case "development":
		c.environment = logger.LogEnvironmentDevelopment
	case "prod":
		fallthrough
	case "production":
		c.environment = logger.LogEnvironmentProduction
	case "":
		return errors.New("environment not provided")
	default:
		return fmt.Errorf("invalid environment string %q", environment)
	}
	return nil
}

func (c *ZapLog) parseLevel(level string) error {
	switch strings.ToLower(level) {
	case "debug":
		c.logLevel = zap.DebugLevel
	case "info":
		c.logLevel = zap.InfoLevel
	case "warn":
		c.logLevel = zap.WarnLevel
	case "error":
		c.logLevel = zap.ErrorLevel
	case "dpanic":
		c.logLevel = zap.DPanicLevel
	case "panic":
		c.logLevel = zap.PanicLevel
	case "fatal":
		c.logLevel = zap.FatalLevel
	case "":
		return errors.New("level not provided")
	default:
		return fmt.Errorf("invalid level string %q", level)
	}
	return nil
}
