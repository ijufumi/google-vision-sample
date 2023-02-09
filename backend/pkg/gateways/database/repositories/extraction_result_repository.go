package repositories

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"github.com/pkg/errors"
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
		return nil, errors.Wrap(err, "ExtractionResultRepository#GetAll")
	}
	return results, nil
}

func (r *extractionResultRepository) GetByID(db *gorm.DB, id string) (*entities.ExtractionResult, error) {
	var result *entities.ExtractionResult
	if err := db.Preload("ExtractedTexts").Where(map[string]string{
		"id": id,
	}).First(result).Error; err != nil {
		return nil, errors.Wrap(err, "ExtractionResultRepository#GetByID")
	}
	return result, nil
}

func (r *extractionResultRepository) Create(db *gorm.DB, entity entities.ExtractionResult) error {
	if err := db.Create(&entity).Error; err != nil {
		return errors.Wrap(err, "ExtractionResultRepository#Create")
	}
	return nil
}

func (r *extractionResultRepository) Update(db *gorm.DB, entity entities.ExtractionResult) error {
	if err := db.Save(&entity).Error; err != nil {
		return errors.Wrap(err, "ExtractionResultRepository#Update")
	}
	return nil
}

func (r *extractionResultRepository) Delete(db *gorm.DB, id string) error {
	if err := db.Model(&entities.ExtractionResult{}).Delete(map[string]string{
		"id": id,
	}).Error; err != nil {
		return errors.Wrap(err, "ExtractionResultRepository#Delete")
	}
	return nil
}
