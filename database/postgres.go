package database

import (
	"fmt"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type PostgresConfig interface {
	ConnectionURL() string
}

func NewPostgresSession(config PostgresConfig) (db.Session, error) {
	if config.ConnectionURL() == "" {
		return nil, fmt.Errorf("invalid connection url")
	}
	parsedUrl, err := postgresql.ParseURL(config.ConnectionURL())
	if err != nil {
		return nil, fmt.Errorf("could not parse Postgres url %q: %w", config.ConnectionURL(), err)
	}
	session, err := postgresql.Open(parsedUrl)
	if err != nil {
		return nil, fmt.Errorf("could not open Postgres session: %w", err)
	}
	return session, nil
}
