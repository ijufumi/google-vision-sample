package repositories

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type JobFileRepository interface {
	GetByJobID(db *gorm.DB, extractionResultID string) ([]*entities.JobFile, error)
	Create(db *gorm.DB, entity ...*entities.JobFile) error
	DeleteByJobID(db *gorm.DB, extractionResultID string) error
}

func NewJobFileRepository() JobFileRepository {
	return &jobFileRepository{}
}

type jobFileRepository struct {
}

func (r *jobFileRepository) GetByJobID(db *gorm.DB, jobID string) ([]*entities.JobFile, error) {
	var results []*entities.JobFile
	if err := db.Where(map[string]string{
		"job_id": jobID,
	}).Find(&results).Error; err != nil {
		return nil, errors.Wrap(err, "JobFileRepository#GetByJobID")
	}
	return results, nil
}

func (r *jobFileRepository) Create(db *gorm.DB, entity ...*entities.JobFile) error {
	if err := db.Create(&entity).Error; err != nil {
		return errors.Wrap(err, "JobFileRepository#Create")
	}
	return nil
}

func (r *jobFileRepository) DeleteByJobID(db *gorm.DB, extractionResultID string) error {
	if err := db.Where("job_id", extractionResultID).Delete(&entities.JobFile{}).Error; err != nil {
		return errors.Wrap(err, "JobFileRepository#DeleteByJobID")
	}
	return nil
}
