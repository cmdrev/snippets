package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() (*viper.Viper, error) {
	viperConfig := viper.New()
	viperConfig.SetConfigFile("config.yml")
	viperConfig.AddConfigPath("./")
	err := viperConfig.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("could not read configuration file: %w", err)
	}
	return viperConfig, nil
}
