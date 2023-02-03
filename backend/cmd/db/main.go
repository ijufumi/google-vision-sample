package main

import (
	"fmt"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/db"
	"github.com/spf13/cobra"
)

func main() {
	config := configs.New()
	migration, err := migrate.New(db.DsnString(config), config.Migration.Path)
	if err != nil {
		panic(err)
	}

	rootCmd := cobra.Command{
		Use: "db",
	}
	rootCmd.AddCommand(upCommand(migration))
	rootCmd.AddCommand(downCommand(migration))
	rootCmd.AddCommand(dropCommand(migration))
}

func upCommand(migration *migrate.Migrate) *cobra.Command {
	return &cobra.Command{
		Use: "up",
		Run: func(cmd *cobra.Command, args []string) {
			err := migration.Up()
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}

func downCommand(migration *migrate.Migrate) *cobra.Command {
	return &cobra.Command{
		Use: "down",
		Run: func(cmd *cobra.Command, args []string) {
			err := migration.Down()
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}

func dropCommand(migration *migrate.Migrate) *cobra.Command {
	return &cobra.Command{
		Use: "drop",
		Run: func(cmd *cobra.Command, args []string) {
			err := migration.Drop()
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}
