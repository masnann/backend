package book_domain

import (
	"gorm.io/gorm"
	"time"
)

type BookModels struct {
	ID        uint64         `json:"id" gorm:"primarykey"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Cover     string         `json:"cover"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
