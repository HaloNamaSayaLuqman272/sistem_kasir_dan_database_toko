package users

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `json:"id" form:"-" gorm:"uuid;primaryKey"`
	Usename     string         `json:"username"`
	Email       string         `json:"email" gorm:"unique"`
	Password    string         `json:"password"`
	PhoneNumber string         `json:"phone_number" gorm:"uniqueIndex"`
	Address     string         `json:"address" gorm:"type:text"`
	ProvinceID  uint           `json:"province_id"`
	CityID      uint           `json:"city_id"`
	DistrictID  uint           `json:"district_id"`
	Role        string         `json:"role"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type EditProfileRequest struct {
	Username       string `form:"username" validate:"required"`
	Email          string `form:"email" validate:"required,email"`
	Password       string `form:"password" validate:"min=8,containsSpecialCharacter,containsNumber"`
	PhoneNumber    string `form:"phone_number" validate:"required,min=10,containsNumberOnly"`
	Address        string `form:"address" validate:"required"`
	ProvinceID     uint   `form:"province_id" validate:"required"`
	CityID         uint   `form:"city_id" validate:"required"`
	DistrictID     uint   `form:"district_id" validate:"required"`
	ProfilePicture string
	File           *multipart.FileHeader
}
