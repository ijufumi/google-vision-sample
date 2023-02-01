package db

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/db"
)

func NewMigration(config *configs.Config) *migrate.Migrate {
	dsn := db.DsnString(config)

	migration, err := migrate.New(fmt.Sprintf("file://%s", config.Migration.Path), dsn)
	if err != nil {
		panic(err)
	}
	return migration
}
