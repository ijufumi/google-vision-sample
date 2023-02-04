package db

import (
	"fmt"

	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *configs.Config) *gorm.DB {
	dsn := DsnString(config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func DsnString(config *configs.Config) string {
	dbConfig := config.DB
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disabl", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
}
