package categories

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CataegoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
