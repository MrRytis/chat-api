package authService

import (
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/repository"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/MrRytis/chat-api/pkg/exception"
	"github.com/google/uuid"
)

func RegisterUser(email string, password string, name string) entity.User {
	hashedPassword := HashPassword(password)
	user := BuildUser(email, hashedPassword, name)
	repository.SaveUser(user)

	return user
}

func FindUserByEmailAndPassword(email string, password string) entity.User {
	user, err := repository.FindUserByEmail(email)
	utils.HandleDbError(err, "User", email)

	if CheckUserPassword(password, user.Password) != nil {
		exception.NewUnauthorized("Invalid credentials")
	}

	return user
}

func LogoutUser(accessToken string, refreshToken string, userId uint, expiresAt int64) {
	BlackListToken(accessToken, expiresAt)
	repository.ExpireRefreshToken(refreshToken, userId)
}

func BuildUser(email string, password string, name string) entity.User {
	return entity.User{
		UUID:     uuid.New().String(),
		Email:    email,
		Password: password,
		Name:     name,
	}
}
