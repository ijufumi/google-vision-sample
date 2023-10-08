package repositories

import (
	"github.com/ijufumi/google-vision-sample/internal/infrastructures/database/entities"
	"gorm.io/gorm"
)

type JobRepository interface {
	GetAll(db *gorm.DB) ([]*entities.Job, error)
	GetByID(db *gorm.DB, id string) (*entities.Job, error)
	Create(db *gorm.DB, entity *entities.Job) error
	Update(db *gorm.DB, entity *entities.Job) error
	Delete(db *gorm.DB, id string) error
}
