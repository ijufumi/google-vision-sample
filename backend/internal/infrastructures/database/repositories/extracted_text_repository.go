package repositories

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ExtractedTextRepository interface {
	GetByID(db *gorm.DB, id string) (*entities.ExtractedText, error)
	Create(db *gorm.DB, entity ...*entities.ExtractedText) error
	DeleteByOutputFileID(db *gorm.DB, outputFileID string) error
}

func NewExtractedTextRepository() ExtractedTextRepository {
	return &extractedTextRepository{}
}

type extractedTextRepository struct {
}

func (r *extractedTextRepository) GetByID(db *gorm.DB, id string) (*entities.ExtractedText, error) {
	var result *entities.ExtractedText
	if err := db.Where("id = ?", id).Find(result).Error; err != nil {
		return nil, errors.Wrap(err, "ExtractedTextRepository#GetByJobID")
	}
	return result, nil
}

func (r *extractedTextRepository) Create(db *gorm.DB, entity ...*entities.ExtractedText) error {
	if len(entity) == 0 {
		return nil
	}
	if err := db.Create(&entity).Error; err != nil {
		return errors.Wrap(err, "ExtractedTextRepository#Create")
	}
	return nil
}

func (r *extractedTextRepository) DeleteByOutputFileID(db *gorm.DB, outputFileID string) error {
	if err := db.Where("output_file_id = ?", outputFileID).Delete(&entities.ExtractedText{}).Error; err != nil {
		return errors.Wrap(err, "ExtractedTextRepository#DeleteByOutputFileID")
	}
	return nil
}
