package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	NameCategory string          `json:"name_category" gorm:"uniqueIndex:idx_name_active,where:deleted_at IS NULL"`
	Description  string          `json:"description"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    *time.Time      `json:"updated_at,omitempty"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
