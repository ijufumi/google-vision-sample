package main

import (
	"fmt"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/db"
	"github.com/ijufumi/google-vision-sample/pkg/loggers"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	config := configs.New()
	logger := loggers.NewLogger()
	sourcePath := fmt.Sprintf("file://%s", config.Migration.Path)
	database := db.NewDB(config, logger.Logger)
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
	rootCmd.AddCommand(createCommand(config.Migration))

	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func createCommand(config configs.Migration) *cobra.Command {
	var ext string
	var name string
	cmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			dir := filepath.Clean(config.Path)
			version := time.Now().UTC().Format("20060102150405")
			for _, direction := range []string{"up", "down"} {
				basename := fmt.Sprintf("%s_%s.%s%s", version, name, direction, ext)
				filename := filepath.Join(dir, basename)

				if err := createFile(filename); err != nil {
					panic(err)
				}

				absPath, _ := filepath.Abs(filename)
				log.Println(absPath)
			}
		},
	}
	cmd.Flags().StringVarP(&ext, "ext", "e", config.Extension, "file extension")
	cmd.Flags().StringVarP(&name, "name", "n", config.Name, "file name")
	return cmd
}

func createFile(filepath string) error {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		return err
	}

	return f.Close()
}

func upCommand(migration *migrate.Migrate) *cobra.Command {
	return &cobra.Command{
		Use: "up",
		Run: func(cmd *cobra.Command, args []string) {
			err := migration.Up()
			if err != nil {
				log.Println(err)
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
				log.Println(err)
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
				log.Println(err)
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
				log.Println(err)
			}
			log.Println(fmt.Sprintf("version is %d", version))
		},
	}
}
