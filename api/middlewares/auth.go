package middlewares

import (
	"sistem_kasir_dan_database_toko/database/models"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	ID   int         `json:"id"`
	Role models.Role `json:"role"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey      string
	ExpireDuration int
}
