package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string `gorm:"unique;not null"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
