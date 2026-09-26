package products

import (
	"context"
	"sistem_kasir_dan_database_toko/package/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type getbycategory struct {
	repository *gorm.DB
}

func (g getbycategory) GetProductsByCategory(ctx context.Context, pagination utils.Pagination, categoryId uint) (utils.Pagination, error) {
	products := []Product{}

	if err := g.repository.WithContext(ctx).Scopes(utils.PaginateByProductCategory(&products, &pagination, categoryId, g.repository)).Preload(clause.Associations).Find(&products).Error; err != nil {
		return utils.Pagination{}, err
	}

	pagination.Rows = products

	return pagination, nil
}
