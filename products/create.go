package products

import (
	"context"
	"sistem_kasir_dan_database_toko/database/models"

	"gorm.io/gorm"
)

type create struct {
	repository *gorm.DB
}

func (c create) CreateProduct(ctx context.Context, createProductRequest *ProductRequest) (Product, error) {
	product := models.Product{
		NameProduct: createProductRequest.NameProduct,
		CategoryID:  createProductRequest.CategoryID,
		Description: createProductRequest.Description,
		Company:     createProductRequest.Company,
		Barcode:     createProductRequest.Barcode,
		Weight:      createProductRequest.Weight,
		ExpiredDate: createProductRequest.ExpiredDate,
		Price:       createProductRequest.Price,
		Stock:       createProductRequest.Stock,
		ImageLink:   createProductRequest.ImageLink,
	}

	result := c.repository.WithContext(ctx).Create(&product)
	if err := result.Error; err != nil {
		return Product{}, err
	}

	record := new(Product)
	if err := result.WithContext(ctx).Preload("Category").Last(record).Error; err != nil {
		return Product{}, nil
	}

	return *record, nil
}
