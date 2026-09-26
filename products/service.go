package products

import (
	"context"
	"sistem_kasir_dan_database_toko/package/ai"
	"sistem_kasir_dan_database_toko/package/utils"

	"gorm.io/gorm"
)

type Service interface {
	CreateProduct(ctx context.Context, createProductRequest *ProductRequest) (Product, error)
	GetProductByID(ctx context.Context, id uint) (Product, error)
	GetProductsByCategory(ctx context.Context, pagination utils.Pagination, categoryId uint) (utils.Pagination, error)
	GetProductsByRecommendation(ctx context.Context, pagination utils.Pagination, promptRequest ai.ProductRecommendationRequest) (utils.Pagination, error)
	GetAllProducts(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error)
	UpdateProductByID(ctx context.Context, updateProductRequest *ProductRequest, id uint) (Product, error)
	DeleteProduct(ctx context.Context, id uint) error
}

type service struct {
	create
	getbyid
	getbycategory
	getbyrecommendation
	getall
	update
	delete
}

var _ Service = (*service)(nil)

func New(repository *gorm.DB) Service {
	return service{
		create:              create{repository: repository},
		getbyid:             getbyid{repository: repository},
		getbycategory:       getbycategory{repository: repository},
		getbyrecommendation: getbyrecommendation{repository: repository},
		getall:              getall{repository: repository},
		update:              update{repository: repository, getbyid: getbyid{repository: repository}},
		delete:              delete{repository: repository, getbyid: getbyid{repository: repository}},
	}
}
