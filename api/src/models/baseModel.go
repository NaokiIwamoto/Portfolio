package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (base *BaseModel) BeforeCreate(db *gorm.DB) error {
	base.CreatedAt = time.Now()
	return nil
}

func (base *BaseModel) BeforeUpdate(db *gorm.DB) error {
	base.UpdatedAt = time.Now()
	return nil
}
