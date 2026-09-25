package categories

import (
	"context"
	"sistem_kasir_dan_database_toko/package/utils"

	"gorm.io/gorm"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAllCategories(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
	categories := []Category{}
	allowedColumns := map[string]bool{
		"name_category": true,
		"description":   true,
	}
	if !allowedColumns[pagination.Keyword] {
		pagination.Keyword = "name_category"
	}

	if err := g.repository.WithContext(ctx).Scopes(utils.Paginate(&categories, &pagination, g.repository)).Find(&categories).Error; err != nil {
		return utils.Pagination{}, err
	}

	pagination.Rows = categories
	return pagination, nil
}
