package services

import (
	"context"
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/db"
	contextManager "github.com/ijufumi/google-vision-sample/pkg/http/context"
	"github.com/ijufumi/google-vision-sample/pkg/loggers"
	"gorm.io/gorm"
)

type baseService struct {
}

func (s *baseService) WithLogger(ctx context.Context, gormDB *gorm.DB, fc func(logger *loggers.Logger, tx *gorm.DB) error) error {
	logger := contextManager.GetLogger(ctx)
	db2 := db.SetLogger(gormDB, logger.Logger)
	return fc(logger, db2)
}
