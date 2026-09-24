package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID      `json:"id" form:"-" gorm:"uuid;primaryKey"`
	Usename        string         `json:"username"`
	Email          string         `json:"email" gorm:"unique"`
	Password       string         `json:"password"`
	PhoneNumber    string         `json:"phone_number"`
	Address        string         `json:"address"`
	ProvinceID     uint           `json:"province_id"`
	CityID         uint           `json:"city_id"`
	DistrictID     uint           `json:"district_id"`
	ProfilePicture string         `json:"profile_picture"`
	Role           string         `json:"role"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

type RegisterRequest struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"min=8,containsSpecialCharacter,containsNumber"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,containsNumberOnly"`
	Address     string `json:"address" validate:"required"`
	ProvinceID  uint   `json:"province_id" validate:"required"`
	CityID      uint   `json:"city_id" validate:"required"`
	DistrictID  uint   `json:"district_id" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
