package db

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func NewMigration() *migrate.Migrate {
	migration, err := migrate.New("", "")
	if err != nil {
		panic(err)
	}
	return migration
}
