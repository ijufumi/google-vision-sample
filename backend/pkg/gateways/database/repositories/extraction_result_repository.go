package repositories

import (
	"github.com/ijufumi/google-vision-sample/pkg/gateways/database/entities"
	"gorm.io/gorm"
)

type ExtractionResultRepository interface {
	GetAll() ([]entities.ExtractionResult, error)
	GetByID(id string) (*entities.ExtractionResult, error)
	Create(entity entities.ExtractionResult) error
	Update(entity entities.ExtractionResult) error
	Delete(id string) error
}

func NewExtractionResultRepository(db *gorm.DB) ExtractionResultRepository {
	return &extractionResultRepository{
		db: db,
	}
}

type extractionResultRepository struct {
	db *gorm.DB
}

func (r *extractionResultRepository) GetAll() ([]entities.ExtractionResult, error) {
	var results []entities.ExtractionResult
	if err := r.db.Preload("ExtractedTexts").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *extractionResultRepository) GetByID(id string) (*entities.ExtractionResult, error) {
	var result *entities.ExtractionResult
	if err := r.db.Preload("ExtractedTexts").Where(map[string]string{
		"id": id,
	}).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *extractionResultRepository) Create(entity entities.ExtractionResult) error {
	if err := r.db.Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *extractionResultRepository) Update(entity entities.ExtractionResult) error {
	if err := r.db.Save(&entity).Error; err != nil {
		return nil
	}
	return nil
}

func (r *extractionResultRepository) Delete(id string) error {
	if err := r.db.Model(&entities.ExtractionResult{}).Delete(map[string]string{
		"id": id,
	}).Error; err != nil {
		return err
	}
	return nil
}
