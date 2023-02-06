package repositories

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"gorm.io/gorm"
)

type ExtractionResultRepository interface {
	GetAll(db *gorm.DB) ([]entities.ExtractionResult, error)
	GetByID(db *gorm.DB, id string) (*entities.ExtractionResult, error)
	Create(db *gorm.DB, entity entities.ExtractionResult) error
	Update(db *gorm.DB, entity entities.ExtractionResult) error
	Delete(db *gorm.DB, id string) error
}

func NewExtractionResultRepository() ExtractionResultRepository {
	return &extractionResultRepository{}
}

type extractionResultRepository struct {
}

func (r *extractionResultRepository) GetAll(db *gorm.DB) ([]entities.ExtractionResult, error) {
	var results []entities.ExtractionResult
	if err := db.Preload("ExtractedTexts").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *extractionResultRepository) GetByID(db *gorm.DB, id string) (*entities.ExtractionResult, error) {
	var result *entities.ExtractionResult
	if err := db.Preload("ExtractedTexts").Where(map[string]string{
		"id": id,
	}).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *extractionResultRepository) Create(db *gorm.DB, entity entities.ExtractionResult) error {
	if err := db.Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *extractionResultRepository) Update(db *gorm.DB, entity entities.ExtractionResult) error {
	if err := db.Save(&entity).Error; err != nil {
		return nil
	}
	return nil
}

func (r *extractionResultRepository) Delete(db *gorm.DB, id string) error {
	if err := db.Model(&entities.ExtractionResult{}).Delete(map[string]string{
		"id": id,
	}).Error; err != nil {
		return err
	}
	return nil
}
