package service

import (
	"context"
	contextManager "github.com/ijufumi/google-vision-sample/internal/common/context"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/db"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type baseService struct {
}

func (s *baseService) Process(ctx context.Context, f func(logger *zap.Logger) error) error {
	logger := contextManager.GetLogger(ctx)
	return f(logger)
}

func (s *baseService) Transaction(gormDB *gorm.DB, logger *zap.Logger, f func(tx *gorm.DB) error) error {
	return gormDB.Transaction(func(tx2 *gorm.DB) error {
		db.SetLogger(tx2, logger)
		return f(tx2)
	})
}

func (s *baseService) Transaction2(ctx context.Context, f func(ctx context.Context) error) error {
	newDB := contextManager.GetDB(ctx)
	return newDB.Transaction(func(tx *gorm.DB) error {
		ctx2 := contextManager.SetTx(ctx, tx)
		return f(ctx2)
	})
}
