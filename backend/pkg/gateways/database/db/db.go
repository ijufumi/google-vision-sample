package db

import (
	"fmt"

	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(config *configs.Config) *gorm.DB {
	dsn := DsnString(config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrate(db)

	return db
}

func DsnString(config *configs.Config) string {
	dbConfig := config.DB
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.ExtractionResults{})
}
