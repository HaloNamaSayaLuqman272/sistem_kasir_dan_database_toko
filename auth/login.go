package auth

import (
	"context"
	"sistem_kasir_dan_database_toko/package/utils"

	"gorm.io/gorm"
)

type login struct {
	repository *gorm.DB
}

func (l login) LoginUser(ctx context.Context, loginRequest *LoginRequest) (User, error) {
	user := new(User)
	if err := l.repository.WithContext(ctx).First(user, "password = ?", loginRequest.Username).Error; err != nil {
		return User{}, err
	}

	if err := utils.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return User{}, err
	}

	return *user, nil
}
