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
	limit := c.QueryInt("limit", 20)
	page := c.QueryInt("page", 1)

	if limit > 100 {
		limit = 100
	}

	if page < 1 {
		page = 1
	}

	groups, total := repository.FindPagedGroupsByUserId(page, limit, c.Locals("userId").(int32))

	var responseGroups []response.Group
	for _, group := range groups {
		var users []response.User

		responseGroups = append(responseGroups, response.Group{
			Uuid: group.Uuid,
			Name: group.Name,
			Admin: response.User{
				Uuid: group.Admin.UUID,
				Name: group.Admin.Name,
			},
			Users: users,
		})
	}

	groupList := response.GroupList{
		PageNumber: int32(page),
		PageSize:   int32(limit),
		ItemsCount: int32(total),
		Items:      responseGroups,
	}

	return c.Status(fiber.StatusOK).JSON(groupList)
}

func GetGroup(c *fiber.Ctx) error {
	groupUuid := c.Params("group")

	group := repository.FindGroupByUuidAndUserId(groupUuid, c.Locals("userId").(int32))

	var users []response.User
	for _, user := range group.Users {
		users = append(users, response.User{
			Uuid: user.UUID,
			Name: user.Name,
		})
	}

	responseGroup := response.Group{
		Uuid: group.Uuid,
		Name: group.Name,
		Admin: response.User{
			Uuid: group.Admin.UUID,
			Name: group.Admin.Name,
		},
		Users: users,
	}

	return c.Status(fiber.StatusOK).JSON(responseGroup)
}

func UpdateGroup(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func DeleteGroup(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func AddUserToGroup(c *fiber.Ctx) error {
	groupUuid := c.Params("group")

	req := new(request.UserToGroup)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	err := utils.ValidateRequest(req)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	group := repository.FindGroupByUuidAndUserId(groupUuid, c.Locals("userId").(int32))
	user := repository.FindUserByUuid(req.Uuid)

	group.Users = append(group.Users, *user)

	repository.UpdateGroup(group)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User added to group",
	})
}

func RemoveUserFromGroup(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}
