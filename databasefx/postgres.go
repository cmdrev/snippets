package databasefx

import (
	"go.uber.org/fx"
	"github.com/cmdrev/snippets/config"
	"github.com/cmdrev/snippets/database"
)

var PostgresModule = fx.Module("postgressession",
	fx.Provide(config.NewPostgresConfig),
	fx.Provide(database.NewPostgresSession),
)
