package categories

import (
	"context"

	"gorm.io/gorm"
)

type getbyid struct {
	repository *gorm.DB
}

func (g getbyid) GetCategoryByID(ctx context.Context, id uint) (Category, error) {
	category := new(Category)
	if err := g.repository.WithContext(ctx).First("id = ?", id).Error; err != nil {
		return Category{}, err
	}

	return *category, nil
}
