package entity

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	CreatedAt time.Time      `gorm:"column:created_at"  json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"  json:"deletedAt"`
}
