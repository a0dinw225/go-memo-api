package models

import (
	"time"
)

type PersonalAccessToken struct {
	ID            uint      `gorm:"primarykey"`
	TokenableType string    `gorm:"column:tokenable_type"`
	TokenableID   uint      `gorm:"column:tokenable_id"`
	Name          string    `gorm:"type:string"`
	Token         string    `gorm:"type:string"`
	Abilities     string    `gorm:"type:text"`
	LastUsedAt    time.Time `gorm:"column:last_used_at"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
