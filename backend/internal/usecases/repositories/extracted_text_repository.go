package repositories

import (
	"context"
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
)

type ExtractedTextRepository interface {
	GetByID(ctx context.Context, id string) (*models.ExtractedText, error)
	Create(ctx context.Context, entity ...*models.ExtractedText) error
	DeleteByOutputFileID(ctx context.Context, outputFileID string) error
}
