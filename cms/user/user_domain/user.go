package user_domain

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID        uint           `json:"id" gorm:"primary-key"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-" gorm:"column:password"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
