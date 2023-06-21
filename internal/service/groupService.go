package service

import (
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/repository"
	"github.com/google/uuid"
)

func BuildGroup(name string, userId int32) entity.Group {
	admin := repository.FindUserById(userId)

	var users []entity.User
	users = append(users, *admin)

	return entity.Group{
		Name:    name,
		Uuid:    uuid.New().String(),
		Admin:   *admin,
		AdminId: int32(admin.ID),
		Users:   users,
	}
}
