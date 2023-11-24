package repositories

import (
	"github.com/ijufumi/google-vision-sample/internal/models/entities"
	"golang.org/x/net/context"
)

type InputFileRepository interface {
	GetByID(ctx context.Context, iD string) (*entities.InputFile, error)
	GetByJobID(ctx context.Context, jobID string) ([]*entities.InputFile, error)
	Create(ctx context.Context, entity ...*entities.InputFile) error
	Update(ctx context.Context, entity *entities.InputFile) error
	DeleteByJobID(ctx context.Context, jobID string) error
}
