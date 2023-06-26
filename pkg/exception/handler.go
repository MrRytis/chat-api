package exception

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError

	response := Exception{
		Code:    status,
		Message: "Internal Server Error",
		Errors:  []Error{},
	}

	var e *Exception
	if errors.As(err, &e) {
		status = e.StatusCode

		response.Code = e.Code
		response.Message = e.Message
		response.Errors = e.Errors
	}

	return c.Status(status).JSON(response)
}
