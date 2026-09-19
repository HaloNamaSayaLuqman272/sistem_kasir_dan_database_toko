package categories

import (
	"context"

	"gorm.io/gorm"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAllCategories(ctx context.Context)
