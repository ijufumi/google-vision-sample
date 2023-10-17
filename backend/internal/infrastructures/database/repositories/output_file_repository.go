package repositories

import (
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
}

func (r *outputFileRepository) GetByJobID(db *gorm.DB, jobID string) ([]*models.OutputFile, error) {
	var results entities.OutputFiles
	if err := db.Where("job_id = ?", jobID).Find(&results).Error; err != nil {
		return nil, errors.Wrap(err, "OutputFileRepository#GetByJobID")
	}
	return results.ToModel(), nil
}

func (r *outputFileRepository) Create(db *gorm.DB, entity ...*models.OutputFile) error {
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
