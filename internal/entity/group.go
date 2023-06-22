package entity

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Users   []User `gorm:"many2many:group_users;"`
	Uuid    string `gorm:"unique"`
	Admin   User   `gorm:"foreignKey:admin_id"`
	AdminId int32  `gorm:"not null"`
}
