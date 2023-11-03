package repositories

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	models "github.com/ijufumi/google-vision-sample/internal/models/entities"
	repositoryInterface "github.com/ijufumi/google-vision-sample/internal/usecases/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewJobRepository() repositoryInterface.JobRepository {
	return &jobRepository{}
}

type jobRepository struct {
}

func (r *jobRepository) GetAll(db *gorm.DB) ([]*models.Job, error) {
	var jobs *entities.Jobs
	if err := db.
		Find(&jobs).Error; err != nil {
		return nil, errors.Wrap(err, "JobRepository#GetAll")
	}
	return jobs.ToModel(), nil
}

func (r *jobRepository) GetByID(db *gorm.DB, id string) (*models.Job, error) {
	var job entities.Job
	if err := db.
		Preload("InputFiles").
		Preload("InputFiles.OutputFiles.ExtractedTexts").
		Where("id = ?", id).First(&job).Error; err != nil {
		return nil, errors.Wrap(err, "JobRepository#GetByID")
	}
	return job.ToModel(), nil
}

func (r *jobRepository) Create(db *gorm.DB, model *models.Job) error {
	entity := entities.FromJobModel(model)
	if err := db.Create(entity).Error; err != nil {
		return errors.Wrap(err, "JobRepository#Create")
	}
	return nil
}

func (r *jobRepository) Update(db *gorm.DB, model *models.Job) error {
	entity := entities.FromJobModel(model)
	if err := db.Save(entity).Error; err != nil {
		return errors.Wrap(err, "JobRepository#Update")
	}
	return nil
}

func (r *jobRepository) Delete(db *gorm.DB, id string) error {
	if err := db.Where("id = ?", id).Delete(&entities.Job{}).Error; err != nil {
		return errors.Wrap(err, "JobRepository#Delete")
	}
	return nil
}
