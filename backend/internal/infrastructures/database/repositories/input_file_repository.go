package repositories

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	repositoryInterface "github.com/ijufumi/google-vision-sample/internal/usecases/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewInputFileRepository() repositoryInterface.InputFileRepository {
	return &inputFileRepository{}
}

type inputFileRepository struct {
}

func (r *inputFileRepository) GetByID(db *gorm.DB, id string) (*entities.InputFile, error) {
	var result entities.InputFile
	if err := db.First(&result, "id = ?", id).Error; err != nil {
		return nil, errors.Wrap(err, "InputFileRepository#GetByID")
	}
	return &result, nil
}

func (r *inputFileRepository) GetByJobID(db *gorm.DB, jobID string) ([]*entities.InputFile, error) {
	var results []*entities.InputFile
	if err := db.Where("job_id = ?", jobID).Find(&results).Error; err != nil {
		return nil, errors.Wrap(err, "InputFileRepository#GetByJobID")
	}
	return results, nil
}

func (r *inputFileRepository) Create(db *gorm.DB, entity ...*entities.InputFile) error {
	if err := db.Create(&entity).Error; err != nil {
		return errors.Wrap(err, "InputFileRepository#Create")
	}
	return nil
}

func (r *inputFileRepository) Update(db *gorm.DB, entity *entities.InputFile) error {
	if err := db.Save(&entity).Error; err != nil {
		return errors.Wrap(err, "InputFileRepository#Update")
	}
	return nil
}

func (r *inputFileRepository) DeleteByJobID(db *gorm.DB, jobID string) error {
	if err := db.Where("job_id = ?", jobID).Delete(&entities.InputFile{}).Error; err != nil {
		return errors.Wrap(err, "InputFileRepository#DeleteByJobID")
	}
	return nil
}
