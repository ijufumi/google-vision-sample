package repositories

import (
	"context"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
	repositoryInterface "github.com/ijufumi/google-vision-sample/internal/usecases/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewExtractedTextRepository() repositoryInterface.ExtractedTextRepository {
	return &extractedTextRepository{}
}

type extractedTextRepository struct {
	baseRepository
}

func (r *extractedTextRepository) GetByID(ctx context.Context, id string) (*models.ExtractedText, error) {
	var result *entities.ExtractedText
	err := r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Find(result).Error; err != nil {
			return errors.Wrap(err, "ExtractedTextRepository#GetByJobID")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result.ToModel(), nil
}

func (r *extractedTextRepository) Create(ctx context.Context, entity ...*models.ExtractedText) error {
	if len(entity) == 0 {
		return nil
	}
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Create(&entity).Error; err != nil {
			return errors.Wrap(err, "ExtractedTextRepository#Create")
		}
		return nil
	})
}

func (r *extractedTextRepository) DeleteByOutputFileID(ctx context.Context, outputFileID string) error {
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Where("output_file_id = ?", outputFileID).Delete(&entities.ExtractedText{}).Error; err != nil {
			return errors.Wrap(err, "ExtractedTextRepository#DeleteByOutputFileID")
		}
		return nil
	})
}
