package entity

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name    string
	Users   []User `gorm:"many2many:user_groups;"`
	Uuid    string `gorm:"unique"`
	Admin   User   `gorm:"foreignKey:admin_id"`
	AdminId uint   `gorm:"not null"`
}
