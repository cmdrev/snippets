package databasefx

import (
	"go.uber.org/fx"
	"shareing/pkg/config"
	"shareing/pkg/database"
)

var PostgresModule = fx.Module("postgressession",
	fx.Provide(config.NewPostgresConfig),
	fx.Provide(database.NewPostgresSession),
)
