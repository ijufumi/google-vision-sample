package repositories

import (
	"context"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
	repositoryInterface "github.com/ijufumi/google-vision-sample/internal/usecases/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewInputFileRepository() repositoryInterface.InputFileRepository {
	return &inputFileRepository{}
}

type inputFileRepository struct {
	baseRepository
}

func (r *inputFileRepository) GetByID(ctx context.Context, id string) (*models.InputFile, error) {
	var result entities.InputFile
	err := r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.First(&result, "id = ?", id).Error; err != nil {
			return errors.Wrap(err, "InputFileRepository#GetByID")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result.ToModel(), nil
}

func (r *inputFileRepository) GetByJobID(ctx context.Context, jobID string) ([]*models.InputFile, error) {
	var results entities.InputFiles
	err := r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Where("job_id = ?", jobID).Find(&results).Error; err != nil {
			return errors.Wrap(err, "InputFileRepository#GetByJobID")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return results.ToModel(), nil
}

func (r *inputFileRepository) Create(ctx context.Context, entity ...*models.InputFile) error {
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Create(&entity).Error; err != nil {
			return errors.Wrap(err, "InputFileRepository#Create")
		}
		return nil
	})
}

func (r *inputFileRepository) Update(ctx context.Context, entity *models.InputFile) error {
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Save(&entity).Error; err != nil {
			return errors.Wrap(err, "InputFileRepository#Update")
		}
		return nil
	})
}

func (r *inputFileRepository) DeleteByJobID(ctx context.Context, jobID string) error {
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Where("job_id = ?", jobID).Delete(&entities.InputFile{}).Error; err != nil {
			return errors.Wrap(err, "InputFileRepository#DeleteByJobID")
		}
		return nil
	})
}
