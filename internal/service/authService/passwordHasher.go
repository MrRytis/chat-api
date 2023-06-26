package authService

import (
	"github.com/MrRytis/chat-api/pkg/exception"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		exception.NewInternalServerError()
	}

	return string(hashedPassword)
}
