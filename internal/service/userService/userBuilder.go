package userService

import (
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/google/uuid"
)

func BuildUser(email string, password string, name string) entity.User {
	return entity.User{
		UUID:     uuid.New().String(),
		Email:    email,
		Password: password,
		Name:     name,
	}
}
