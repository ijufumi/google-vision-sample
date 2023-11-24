package repositories

import (
	"context"
	"github.com/ijufumi/google-vision-sample/internal/models/entities"
)

type OutputFileRepository interface {
	GetByJobID(ctx context.Context, jobID string) ([]*entities.OutputFile, error)
	Create(ctx context.Context, entity ...*entities.OutputFile) error
	DeleteByJobID(ctx context.Context, jobID string) error
	DeleteByInputFileID(ctx context.Context, inputFileID string) error
}
