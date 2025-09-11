package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	ginHttpData struct {
		Address string `yaml:"address"`
	}

	GinHttp struct {
		Http
		data ginHttpData
	}
)

func NewGinHttpConfig(viper *viper.Viper) (*GinHttp, error) {
	httpConfig := GinHttp{data: ginHttpData{}}
	err := viper.UnmarshalKey("http", &httpConfig.data)
	if err != nil {
		return &GinHttp{}, fmt.Errorf("could not unmarshal gin http configuration: %w", err)
	}
	return &httpConfig, nil
}

func (p *GinHttp) Address() string {
	return p.data.Address
}
