package products

import (
	"context"
	"sistem_kasir_dan_database_toko/package/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAllProducts(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
	products := []Product{}
	if err := g.repository.WithContext(ctx).Scopes(utils.Paginate(&products, &pagination, g.repository)).Preload(clause.Associations).Find(&products).Error; err != nil {
		return utils.Pagination{}, err
	}

	pagination.Rows = products
	return pagination, nil
}
