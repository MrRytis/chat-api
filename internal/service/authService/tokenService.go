package authService

import (
	"fmt"
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/repository"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/MrRytis/chat-api/pkg/exception"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

var AccessTokenJwtExpDuration = time.Hour * 2

func CreateAccessToken(user entity.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":      user.UUID,
		"uid":       user.ID,
		"name":      user.Name,
		"expiresAt": time.Now().Add(AccessTokenJwtExpDuration).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		exception.NewInternalServerError()
	}

	return tokenString
}

func CreateRefreshToken(user entity.User) string {
	u := uuid.New().String()

	token := entity.RefreshToken{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
		UserId:    user.ID,
		Token:     u,
	}

	repository.CreateRefreshToken(token)

	return token.Token
}

func RefreshAccessToken(refreshToken string, accessToken string) string {
	rt, err := repository.FindRefreshTokenByToken(refreshToken)
	if err != nil {
		exception.NewInternalServerError()
	}

	claims, err := ParseJWT(accessToken)
	if err != nil {
		exception.NewInternalServerError()
	}

	if rt.UserId != claims["uuid"].(uint) {
		exception.NewInternalServerError()
	}

	newAccessToken := CreateAccessToken(rt.User)

	return newAccessToken
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func BlackListToken(token string, expiresAt int64) {
	utils.SetCache(blackListKey(token), token, time.Unix(expiresAt, 0).Sub(time.Now()))
}

func IsBlacklisted(token string) bool {
	var value bool
	val := utils.GetFromCache(blackListKey(token), value)
	if val != nil {
		return true
	}

	return false
}

func blackListKey(token string) string {
	return fmt.Sprintf("T-%s", token)
}
