package repositories

import (
	"context"
	contextManager "github.com/ijufumi/google-vision-sample/internal/common/context"
	"gorm.io/gorm"
)

type baseRepository struct{}

func (r *baseRepository) Transaction(ctx context.Context, f func(tx *gorm.DB) error) error {
	tx := contextManager.GetDB(ctx)
	return f(tx)
}
