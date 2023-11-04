package service

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/db"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type baseService struct {
}

func (s *baseService) Transaction(gormDB *gorm.DB, logger *zap.Logger, f func(tx *gorm.DB) error) error {
	return gormDB.Transaction(func(tx2 *gorm.DB) error {
		db.SetLogger(tx2, logger)
		return f(tx2)
	})
}
