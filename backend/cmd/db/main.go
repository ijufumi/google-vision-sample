package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/db"
	"github.com/spf13/cobra"
)

func main() {
	config := configs.New()
	sourcePath := fmt.Sprintf("file://%s", config.Migration.Path)
	database := db.NewDB(config)
	sqlDB, err := database.DB()
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	migration, err := migrate.NewWithDatabaseInstance(sourcePath, config.DB.Name, driver)
	if err != nil {
		panic(err)
	}

	rootCmd := &cobra.Command{
		Use: "db",
	}
	rootCmd.AddCommand(upCommand(migration))
	rootCmd.AddCommand(downCommand(migration))
	rootCmd.AddCommand(dropCommand(migration))
	rootCmd.AddCommand(versionCommand(migration))
	rootCmd.AddCommand(createCommand())

	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func createCommand() *cobra.Command {
	return &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
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

func versionCommand(migration *migrate.Migrate) *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			version, _, err := migration.Version()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(fmt.Sprintf("version is %d", version))
		},
	}
}
