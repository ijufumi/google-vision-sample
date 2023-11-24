package repositories

import (
	"context"
	contextManager "github.com/ijufumi/google-vision-sample/internal/common/context"
	"gorm.io/gorm"
)

type baseRepository struct{}

func (r *baseRepository) Transaction(ctx context.Context, f func(tx *gorm.DB) error) error {
	if contextManager.HasTx(ctx) {
		tx := contextManager.GetTx(ctx)
		return f(tx)
	}
	db := contextManager.GetDB(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		return f(tx)
	})
}
