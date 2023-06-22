package middleware

import (
	"github.com/MrRytis/chat-api/internal/service/authService"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Auth(c *fiber.Ctx) error {
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	if tokenString == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
	}

	claims, err := authService.ParseJWT(tokenString)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// Check if the token is blacklisted (User logged out)
	if authService.IsBlacklisted(tokenString) {
		return fiber.NewError(fiber.StatusUnauthorized, "Token is blacklisted")
	}

	uuid := claims["uuid"].(string)
	userId := int32(claims["uid"].(float64))
	expiresAt := int64(claims["expiresAt"].(float64))

	c.Locals("uuid", uuid)
	c.Locals("userId", userId)
	c.Locals("expiresAt", expiresAt)
	c.Locals("jwt", tokenString)

	return c.Next()
}
