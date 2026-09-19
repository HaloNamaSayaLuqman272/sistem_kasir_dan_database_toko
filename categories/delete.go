package categories

import (
	"context"

	"gorm.io/gorm"
)

type deletebyid struct {
	repository *gorm.DB
	getbyid
}

func (d deletebyid) DeleteCategoryByID(ctx context.Context, id uint) error {
	category, err := d.GetCategoryByID(ctx, id)
	if err != nil {
		return err
	}

	if err := d.repository.WithContext(ctx).Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
