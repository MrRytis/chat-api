package middleware

import (
	"github.com/MrRytis/chat-api/pkg/exception"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WebsocketProtocol(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}

	exception.NewUpgradeRequired("Upgrade required")
	return nil
}
