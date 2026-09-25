package categories

import (
	"context"
	"sistem_kasir_dan_database_toko/database/models"

	"gorm.io/gorm"
)

type create struct {
	repository *gorm.DB
}

func (c create) CreateCategory(ctx context.Context, createCategoryRequest *CataegoryRequest) (Category, error) {
	category := models.Category{
		NameCategory: createCategoryRequest.NameCategory,
		Description:  createCategoryRequest.Description,
	}
	result := c.repository.WithContext(ctx).Create(&category)
	record := new(Category)
	if err := result.Error; err != nil {
		return Category{}, err
	}

	if err := result.WithContext(ctx).Last(record).Error; err != nil {
		return Category{}, err
	}

	return *record, nil
}
