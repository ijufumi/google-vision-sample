package repositories

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ExtractedTextRepository interface {
	GetByExtractionResultID(db *gorm.DB, extractionResultID string) ([]*entities.ExtractedText, error)
	Create(db *gorm.DB, entity ...*entities.ExtractedText) error
	DeleteByOutputFileID(db *gorm.DB, outputFileID string) error
}

func NewExtractedTextRepository() ExtractedTextRepository {
	return &extractedTextRepository{}
}

type extractedTextRepository struct {
}

func (r *extractedTextRepository) GetByExtractionResultID(db *gorm.DB, extractionResultID string) ([]*entities.ExtractedText, error) {
	var results []*entities.ExtractedText
	if err := db.Where(map[string]string{
		"extractionResultID": extractionResultID,
	}).Find(&results).Error; err != nil {
		return nil, errors.Wrap(err, "ExtractedTextRepository#GetByJobID")
	}
	return results, nil
}

func (r *extractedTextRepository) Create(db *gorm.DB, entity ...*entities.ExtractedText) error {
	if err := db.Create(&entity).Error; err != nil {
		return errors.Wrap(err, "ExtractedTextRepository#Create")
	}
	return nil
}

func (r *extractedTextRepository) DeleteByOutputFileID(db *gorm.DB, outputFileID string) error {
	if err := db.Where("output_file_id", outputFileID).Delete(&entities.ExtractedText{}).Error; err != nil {
		return errors.Wrap(err, "ExtractedTextRepository#DeleteByOutputFileID")
	}
	return nil
}
