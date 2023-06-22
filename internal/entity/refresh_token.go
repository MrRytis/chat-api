package entity

import (
	"time"
)

type RefreshToken struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
	UserId    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:user_id"`
	Token     string `gorm:"not null"`
}
