package api

import (
	"sistem_kasir_dan_database_toko/api/handlers"
	"sistem_kasir_dan_database_toko/api/middlewares"
	"sistem_kasir_dan_database_toko/auth"
	"sistem_kasir_dan_database_toko/categories"
	"sistem_kasir_dan_database_toko/package/fileupload"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func NewEcho(repository *gorm.DB, cld *cloudinary.Cloudinary, jwtConfig middlewares.JWTConfig) *echo.Echo {
	var (
		e                 = echo.New()
		authService       = auth.New(repository)
		categoryService   = categories.New(repository)
		uploader          = &fileupload.CloudinaryUploader{Cld: cld}
		authHandler       = handlers.NewAuth(authService, jwtConfig)
		categoriesHandler = handlers.NewCategories(categoryService)
	)
}
