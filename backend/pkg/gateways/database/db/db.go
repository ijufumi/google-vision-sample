package db

import (
	"fmt"

	"github.com/ijufumi/google-vision-sample/pkg/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(config *configs.Config) *gorm.DB {
	dbConfig := config.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
