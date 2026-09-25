package middlewares

import (
	"context"
	"errors"
	"net/http"
	"sistem_kasir_dan_database_toko/database/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

type JWTCustomClaims struct {
	ID   uuid.UUID   `json:"id"`
	Role models.Role `json:"role"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey      string
	ExpireDuration int
}

type ContextKey string

const userContextKey = ContextKey("user")

func (jwtConfig *JWTConfig) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(JWTCustomClaims)
		},
		SigningKey: []byte(jwtConfig.SecretKey),
	}
}

func (jwtConfig *JWTConfig) GenerateToken(userID uuid.UUID, role models.Role) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(jwtConfig.ExpireDuration)))

	claims := &JWTCustomClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
		Role: role,
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := rawToken.SignedString([]byte(jwtConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(ctx context.Context) (*JWTCustomClaims, error) {
	user, ok := ctx.Value(userContextKey).(*jwt.Token)
	if !ok || user == nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := user.Claims.(*JWTCustomClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	claims, err := GetUser(ctx)
	if err != nil {
		return uuid.Nil, errors.New("invalid token")
	}

	return claims.ID, nil
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		ctx := context.WithValue(c.Request().Context(), userContextKey, user)
		c.SetRequest(c.Request().WithContext(ctx))

		userData, err := GetUser(ctx)
		if userData == nil || err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		c.Set("userData", userData)
		return next(c)
	}
}

func VerifyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		user, err := GetUser(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		if user.Role != models.Admin {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "access denied",
			})
		}

		return next(c)
	}
}
