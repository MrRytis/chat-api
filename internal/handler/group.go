package handler

import (
	"github.com/MrRytis/chat-api/internal/model/request"
	"github.com/MrRytis/chat-api/internal/model/response"
	"github.com/MrRytis/chat-api/internal/repository"
	"github.com/MrRytis/chat-api/internal/service"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateGroup(c *fiber.Ctx) error {
	req := new(request.CreateGroup)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	err := utils.ValidateRequest(req)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	group := service.BuildGroup(req.Name, c.Locals("userId").(int32))
	repository.CreateGroup(group)

	var users []response.User
	for _, user := range group.Users {
		users = append(users, response.User{
			Uuid: user.UUID,
			Name: user.Name,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.Group{
		Uuid: group.Uuid,
		Name: group.Name,
		Admin: response.User{
			Uuid: group.Admin.UUID,
			Name: group.Admin.Name,
		},
		Users: users,
	})
}

func GetGroupList(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func GetGroup(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func UpdateGroup(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func DeleteGroup(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func AddUserToGroup(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func RemoveUserFromGroup(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}
