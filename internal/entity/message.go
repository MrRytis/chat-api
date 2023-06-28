package entity

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UUID     string `gorm:"unique;not null"`
	UserId   uint   `gorm:"not null"`
	User     User   `gorm:"foreignKey:user_id"`
	GroupId  uint   `gorm:"not null"`
	Group    Group  `gorm:"foreignKey:group_id"`
	Content  string `gorm:"not null"`
	IsEdited bool   `gorm:"not null"`
}
