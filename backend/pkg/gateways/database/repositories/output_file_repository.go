package repositories

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type OutputFileRepository interface {
	GetByJobID(db *gorm.DB, jobID string) ([]*entities.OutputFile, error)
	Create(db *gorm.DB, entity ...*entities.OutputFile) error
	DeleteByJobID(db *gorm.DB, jobID string) error
	DeleteByInputFileID(db *gorm.DB, inputFileID string) error
}

func NewOutputFileRepository() OutputFileRepository {
	return &outputFileRepository{}
}

type outputFileRepository struct {
}

func (r *outputFileRepository) GetByJobID(db *gorm.DB, jobID string) ([]*entities.OutputFile, error) {
	var results []*entities.OutputFile
	if err := db.Where("job_id = ?", jobID).Find(&results).Error; err != nil {
		return nil, errors.Wrap(err, "OutputFileRepository#GetByJobID")
	}
	return results, nil
}

func (r *outputFileRepository) Create(db *gorm.DB, entity ...*entities.OutputFile) error {
	if err := db.Create(&entity).Error; err != nil {
		return errors.Wrap(err, "OutputFileRepository#Create")
	}
	return nil
}

func (r *outputFileRepository) DeleteByJobID(db *gorm.DB, jobID string) error {
	if err := db.Where("job_id = ?", jobID).Delete(&entities.OutputFile{}).Error; err != nil {
		return errors.Wrap(err, "OutputFileRepository#DeleteByJobID")
	}
	return nil
}

func (r *outputFileRepository) DeleteByInputFileID(db *gorm.DB, inputFileID string) error {
	if err := db.Where("input_file_id = ?", inputFileID).Delete(&entities.OutputFile{}).Error; err != nil {
		return errors.Wrap(err, "OutputFileRepository#DeleteByJobID")
	}
	return nil
}
