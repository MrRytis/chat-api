package repository

import (
	"errors"
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/MrRytis/chat-api/pkg/exception"
	"gorm.io/gorm"
	"log"
	"time"
)

func CreateRefreshToken(refreshToken entity.RefreshToken) entity.RefreshToken {
	err := utils.Db.Create(&refreshToken).Error
	if err != nil {
		exception.NewInternalServerError()
	}

	return refreshToken
}

func ExpireRefreshToken(token string, userId uint) {
	var refreshToken entity.RefreshToken

	err := utils.Db.Where("token = ? AND user_id = ? AND expires_at > ?", token, userId, time.Now()).First(&refreshToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}

		log.Fatal(err, "Error finding refresh token")
	}

	refreshToken.ExpiresAt = time.Now()

	err = utils.Db.Save(&refreshToken).Error
	if err != nil {
		log.Fatal(err, "Error saving refresh token")
	}
}

func FindRefreshTokenByToken(token string) (entity.RefreshToken, error) {
	var refreshToken entity.RefreshToken

	err := utils.Db.Where("token = ? AND expires_at > ?", token, time.Now()).First(&refreshToken).Error
	if err != nil {
		return entity.RefreshToken{}, err
	}

	return refreshToken, nil
}
