package handler

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Error()
	}

	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": message,
	})
}
