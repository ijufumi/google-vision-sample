package repositories

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"gorm.io/gorm"
)

type ExtractedTextRepository interface {
	GetByExtractionResultID(extractionResultID string) ([]entities.ExtractedText, error)
	Create(entity entities.ExtractedText) error
	DeleteByExtractionResultID(extractionResultID string) error
}

func NewExtractedTextRepository(db *gorm.DB) ExtractedTextRepository {
	return &extractedTextRepository{
		db: db,
	}
}

type extractedTextRepository struct {
	db *gorm.DB
}

func (r *extractedTextRepository) GetByExtractionResultID(extractionResultID string) ([]entities.ExtractedText, error) {
	var results []entities.ExtractedText
	if err := r.db.Where(map[string]string{
		"extractionResultID": extractionResultID,
	}).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *extractedTextRepository) Create(entity entities.ExtractedText) error {
	if err := r.db.Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *extractedTextRepository) DeleteByExtractionResultID(extractionResultID string) error {
	if err := r.db.Model(&entities.ExtractedText{}).Delete(map[string]string{
		"extractionResultID": extractionResultID,
	}).Error; err != nil {
		return err
	}
	return nil
}
