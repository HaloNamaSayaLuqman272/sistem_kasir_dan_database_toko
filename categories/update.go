package categories

import (
	"context"
	"sistem_kasir_dan_database_toko/database/models"

	"gorm.io/gorm"
)

type updatecategory struct {
	repository *gorm.DB
	getbyid
}

func (u updatecategory) UpdateCategoryByID(ctx context.Context, updateCategoryRequest *CataegoryRequest, id uint) (Category, error) {
	category := models.Category{
		NameCategory: updateCategoryRequest.NameCategory,
		Description:  updateCategoryRequest.Description,
	}
	if err := u.repository.WithContext(ctx).Where("id = ?", id).Updates(&category).Error; err != nil {
		return Category{}, err
	}

	record, err := u.getbyid.GetCategoryByID(ctx, id)
	if err != nil {
		return Category{}, err
	}

	return record, nil
}
