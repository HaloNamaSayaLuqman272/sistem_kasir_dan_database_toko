package models

import (
	"database/sql/driver"
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
	Role        Role           `json:"role"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Role string

const (
	Enduser Role = "user"
	Admin   Role = "admin"
)

func (p *Role) Scan(value any) error {
	*p = Role(value.([]byte))
	return nil
}

func (p Role) Value() (driver.Value, error) {
	return string(p), nil
}
