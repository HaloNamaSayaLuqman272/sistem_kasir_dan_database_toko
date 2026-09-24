package main

import (
	"log"
	"sistem_kasir_dan_database_toko/api/middlewares"
	"sistem_kasir_dan_database_toko/database/drivers"
	"sistem_kasir_dan_database_toko/package/constant"
	"sistem_kasir_dan_database_toko/package/fileupload"
	"sistem_kasir_dan_database_toko/package/utils"
	"strconv"
)

func main() {
	dbConfig := drivers.DBConfig{
		Username: utils.GetConfigurance(constant.DB_USERNAME),
		Password: utils.GetConfigurance(constant.DB_PASSWORD),
		Database: utils.GetConfigurance(constant.DB_NAME),
		Host:     utils.GetConfigurance(constant.DB_HOST),
		Port:     utils.GetConfigurance(constant.DB_PORT),
	}

	clurdinaryConfig := fileupload.CloudinaryConfig{
		CloudinaryURL: utils.GetConfigurance(constant.CLOUDINARY_URL),
	}

	expireDuration, err := strconv.Atoi(utils.GetConfigurance(constant.JWT_EXPIRE_DURATION))
	if err != nil {
		log.Fatalf("error when parsing expire duration: %v\n", err)
	}

	jwtConfig := middlewares.JWTConfig{
		SecretKey:      utils.GetConfigurance(constant.JWT_SECRET_KEY),
		ExpireDuration: expireDuration,
	}

	var (
		repository = dbConfig.InitDB()
		cloudinary = clurdinaryConfig.InitCloudinary()
	)

	drivers.MigrateDB(repository)
}
