package config

import (
	"fmt"
	"github.com/spf13/viper"
	"shareing/pkg/database"
)

type postgresData struct {
	ConnectionURL string `yaml:"connectionURL"`
}

type Postgres struct {
	database.PostgresConfig
	postgresData
}

func NewPostgresConfig(viper *viper.Viper) (database.PostgresConfig, error) {
	postgresConfig := Postgres{}
	postgresConfig.postgresData = postgresData{}
	err := viper.UnmarshalKey("database.postgres", &postgresConfig.postgresData)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal postgres configuration: %w", err)
	}
	return &postgresConfig, nil
}

func (p *Postgres) ConnectionURL() string {
	return p.postgresData.ConnectionURL
}
