package repositories

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type JobRepository interface {
	GetAll(db *gorm.DB) ([]*entities.Job, error)
	GetByID(db *gorm.DB, id string) (*entities.Job, error)
	Create(db *gorm.DB, entity *entities.Job) error
	Update(db *gorm.DB, entity *entities.Job) error
	Delete(db *gorm.DB, id string) error
}

func NewJobRepository() JobRepository {
	return &jobRepository{}
}

type jobRepository struct {
}

func (r *jobRepository) GetAll(db *gorm.DB) ([]*entities.Job, error) {
	var results []*entities.Job
	if err := db.
		Find(&results).Error; err != nil {
		return nil, errors.Wrap(err, "JobRepository#GetAll")
	}
	return results, nil
}

func (r *jobRepository) GetByID(db *gorm.DB, id string) (*entities.Job, error) {
	var result entities.Job
	if err := db.
		Preload("InputFiles").
		Preload("InputFiles.OutputFiles").
		Preload("InputFiles.OutputFiles.ExtractedTexts").
		Where("id = ?", id).First(&result).Error; err != nil {
		return nil, errors.Wrap(err, "JobRepository#GetByID")
	}
	return &result, nil
}

func (r *jobRepository) Create(db *gorm.DB, entity *entities.Job) error {
	if err := db.Create(entity).Error; err != nil {
		return errors.Wrap(err, "JobRepository#Create")
	}
	return nil
}

func (r *jobRepository) Update(db *gorm.DB, entity *entities.Job) error {
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
