package config

import (
	"fmt"
	"github.com/cmdrev/snippets/database"
	"github.com/spf13/viper"
)

type mongoDBData struct {
	ConnectionURI string `yaml:"connectionURI"`
}

type MongoDB struct {
	database.MongoDBConfig
	mongoDBData
}

func NewMongoDBConfig(viper *viper.Viper) (database.MongoDBConfig, error) {
	mongoDBConfig := MongoDB{}
	mongoDBConfig.mongoDBData = mongoDBData{}
	err := viper.UnmarshalKey("database.mongodb", &mongoDBConfig.mongoDBData)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal MongoDB configuration: %w", err)
	}
	return &mongoDBConfig, nil
}

func (p *MongoDB) ConnectionURI() string {
	return p.mongoDBData.ConnectionURI
}
