package repositories

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
	repositoryInterface "github.com/ijufumi/google-vision-sample/internal/usecases/repositories"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func NewJobRepository() repositoryInterface.JobRepository {
	return &jobRepository{}
}

type jobRepository struct {
	baseRepository
}

func (r *jobRepository) GetAll(ctx context.Context) ([]*models.Job, error) {
	var jobs *entities.Jobs
	err := r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.
			Find(&jobs).Error; err != nil {
			return errors.Wrap(err, "JobRepository#GetAll")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return jobs.ToModel(), nil
}

func (r *jobRepository) GetByID(ctx context.Context, id string) (*models.Job, error) {
	var job entities.Job
	err := r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.
			Preload("InputFiles").
			Preload("InputFiles.OutputFiles.ExtractedTexts").
			Where("id = ?", id).First(&job).Error; err != nil {
			return errors.Wrap(err, "JobRepository#GetByID")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return job.ToModel(), nil
}

func (r *jobRepository) Create(ctx context.Context, model *models.Job) error {
	entity := entities.FromJobModel(model)
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Create(entity).Error; err != nil {
			return errors.Wrap(err, "JobRepository#Create")
		}
		return nil
	})
}

func (r *jobRepository) Update(ctx context.Context, model *models.Job) error {
	entity := entities.FromJobModel(model)
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Save(entity).Error; err != nil {
			return errors.Wrap(err, "JobRepository#Update")
		}
		return nil
	})
}

func (r *jobRepository) Delete(ctx context.Context, id string) error {
	return r.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&entities.Job{}).Error; err != nil {
			return errors.Wrap(err, "JobRepository#Delete")
		}
		return nil
	})
}
