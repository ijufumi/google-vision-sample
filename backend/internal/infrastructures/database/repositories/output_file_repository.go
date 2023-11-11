package repositories

import (
	"context"
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
	repositoryInterface "github.com/ijufumi/google-vision-sample/internal/usecases/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewOutputFileRepository() repositoryInterface.OutputFileRepository {
	return &outputFileRepository{}
}

type outputFileRepository struct {
	baseRepository
}

func (r *outputFileRepository) GetByJobID(ctx context.Context, jobID string) ([]*models.OutputFile, error) {
	var results entities.OutputFiles
	err := r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Where("job_id = ?", jobID).Find(&results).Error; err != nil {
			return errors.Wrap(err, "OutputFileRepository#GetByJobID")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return results.ToModel(), nil
}

func (r *outputFileRepository) Create(ctx context.Context, entity ...*models.OutputFile) error {
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Create(&entity).Error; err != nil {
			return errors.Wrap(err, "OutputFileRepository#Create")
		}
		return nil
	})
}

func (r *outputFileRepository) DeleteByJobID(ctx context.Context, jobID string) error {
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Where("job_id = ?", jobID).Delete(&entities.OutputFile{}).Error; err != nil {
			return errors.Wrap(err, "OutputFileRepository#DeleteByJobID")
		}
		return nil
	})
}

func (r *outputFileRepository) DeleteByInputFileID(ctx context.Context, inputFileID string) error {
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Where("input_file_id = ?", inputFileID).Delete(&entities.OutputFile{}).Error; err != nil {
			return errors.Wrap(err, "OutputFileRepository#DeleteByJobID")
		}
		return nil
	})
}
