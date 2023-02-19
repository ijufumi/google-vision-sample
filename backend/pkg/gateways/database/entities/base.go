package entities

import (
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	gorm.Model
}

func (e *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	e.CreatedAt = time.Now().UTC()
	return
}

func (e *BaseEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	e.UpdatedAt = time.Now().UTC()
	return
}
