package repository

import (
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/MrRytis/chat-api/pkg/exception"
)

func FindUserById(id int32) (entity.User, error) {
	var user entity.User

	err := utils.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func FindUserByEmail(email string) (entity.User, error) {
	var user entity.User

	err := utils.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func SaveUser(user entity.User) {
	err := utils.Db.Create(&user).Error
	if err != nil {
		exception.NewInternalServerError()
	}
}

func FindUserByUuid(uuid string) (entity.User, error) {
	var users entity.User

	err := utils.Db.Where("uuid = ?", uuid).First(&users).Error
	if err != nil {
		return entity.User{}, err
	}

	return users, nil
}
