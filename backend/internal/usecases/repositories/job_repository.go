package repositories

import (
	"context"
	"github.com/ijufumi/google-vision-sample/internal/models/entities"
)

type JobRepository interface {
	GetAll(ctx context.Context) ([]*entities.Job, error)
	GetByID(ctx context.Context, id string) (*entities.Job, error)
	Create(ctx context.Context, entity *entities.Job) error
	Update(ctx context.Context, entity *entities.Job) error
	Delete(ctx context.Context, id string) error
}
