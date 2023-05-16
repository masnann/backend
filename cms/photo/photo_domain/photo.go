package photo_domain

import (
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	ID         uint64         `json:"id" gorm:"primarykey"`
	Images     string         `json:"images"`
	CategoryID uint           `json:"category_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Category struct {
	ID        uint64         `json:"id" gorm:"primarykey"`
	Name      string         `json:"name"`
	Photos    []Photo        `json:"photos"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
