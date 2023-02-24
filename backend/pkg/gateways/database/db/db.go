package db

import (
	"fmt"
	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func NewDB(config *configs.Config, zapLogger *zap.Logger) *gorm.DB {
	dsn := dsnString(config)
	logger := zapgorm2.New(zapLogger)
	logger.SetAsDefault()
	logger.LogMode(gormLogger.Info)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		panic(err)
	}

	return db
}

func dsnString(config *configs.Config) string {
	dbConfig := config.DB
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
}
