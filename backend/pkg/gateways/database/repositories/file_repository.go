package repositories

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FileRepository interface {
	GetByExtractionResultID(db *gorm.DB, extractionResultID string) ([]*entities.File, error)
	Create(db *gorm.DB, entity ...*entities.File) error
	DeleteByExtractionResultID(db *gorm.DB, extractionResultID string) error
}

func NewFileRepository() FileRepository {
	return &fileRepository{}
}

type fileRepository struct {
}

func (r *fileRepository) GetByExtractionResultID(db *gorm.DB, extractionResultID string) ([]*entities.File, error) {
	var results []*entities.File
	if err := db.Where(map[string]string{
		"extractionResultID": extractionResultID,
	}).Find(&results).Error; err != nil {
		return nil, errors.Wrap(err, "FileRepository#GetByExtractionResultID")
	}
	return results, nil
}

func (r *fileRepository) Create(db *gorm.DB, entity ...*entities.File) error {
	if err := db.Create(&entity).Error; err != nil {
		return errors.Wrap(err, "FileRepository#Create")
	}
	return nil
}

func (r *fileRepository) DeleteByExtractionResultID(db *gorm.DB, extractionResultID string) error {
	if err := db.Where("extraction_result_id", extractionResultID).Delete(&entities.File{}).Error; err != nil {
		return errors.Wrap(err, "FileRepository#DeleteByExtractionResultID")
	}
	return nil
}
