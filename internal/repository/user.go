package repository

import (
	"errors"
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/utils"
	"gorm.io/gorm"
	"log"
)

func FindUserById(id int32) *entity.User {
	var user entity.User

	err := utils.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		log.Fatal(err, "Error finding user")
	}

	return &user
}

func FindUserByEmail(email string) *entity.User {
	var user entity.User

	err := utils.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		log.Fatal(err, "Error finding user")
	}

	return &user
}

func SaveUser(user entity.User) {
	err := utils.Db.Create(&user).Error
	if err != nil {
		log.Fatal(err, "Error saving user")
	}
}

func FindUserByUuid(uuid string) *entity.User {
	var users *entity.User

	err := utils.Db.Where("uuid = ?", uuid).First(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		log.Fatal(err, "Error finding users")
	}

	return users
}
