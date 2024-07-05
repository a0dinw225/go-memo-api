package models

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        int            `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	UserID    int            `json:"user_id"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
