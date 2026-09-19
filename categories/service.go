package categories

import (
	"context"
	"sistem_kasir_dan_database_toko/package/utils"

	"gorm.io/gorm"
)

type Service interface {
	CreateCategory(ctx context.Context, createCategoryRequest *CataegoryRequest) (Category, error)
	GetCategoryByID(ctx context.Context, id uint) (Category, error)
	GetAllCategories(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error)
	UpdateCategoryByID(ctx context.Context, updateCategoryRequest *CataegoryRequest, id uint) (Category, error)
	DeleteCategoryByID(ctx context.Context, id uint) error
}

type service struct {
	create
	getbyid
	getall
	updatecategory
	deletebyid
}

var _ Service = (*service)(nil)

func New(repository *gorm.DB) Service {
	return service{
		create:         create{repository: repository},
		getbyid:        getbyid{repository: repository},
		getall:         getall{repository: repository},
		updatecategory: updatecategory{repository: repository, getbyid: getbyid{repository: repository}},
		deletebyid:     deletebyid{repository: repository, getbyid: getbyid{repository: repository}},
	}
}
