package handler

import (
	"github.com/MrRytis/chat-api/internal/model/request"
	"github.com/MrRytis/chat-api/internal/model/response"
	"github.com/MrRytis/chat-api/internal/repository"
	"github.com/MrRytis/chat-api/internal/service"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateGroup godoc
// @Summary      Create new group
// @Description  creates new group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        req body request.CreateGroup true "create group"
// @Success      201  {object}  response.Group
// @Failure      400  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups [post]
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

// GetGroupList godoc
// @Summary      Get group list
// @Description  Get paginated group list
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        page query int false "default 1"
// @Param        limit query int false "default 20"
// @Success      200  {object}  response.GroupList
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups [get]
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

// GetGroup godoc
// @Summary      Get single group
// @Description  get single group by uuid
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        uuid path string true "uuid of the group"
// @Success      200  {object}  response.Group
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{uuid} [get]
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

// UpdateGroup godoc
// @Summary      Update group
// @Description  update group by uuid
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        req body request.CreateGroup true "update group"
// @Param        uuid path string true "uuid of the group"
// @Success      200  {object}  response.Group
// @Failure      400  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{uuid} [put]
func UpdateGroup(c *fiber.Ctx) error {
	groupUuid := c.Params("group")

	req := new(request.CreateGroup)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse JSON body")
	}

	err := utils.ValidateRequest(req)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	group := repository.FindGroupByUuidAndUserId(groupUuid, c.Locals("userId").(int32))

	if group.AdminId != c.Locals("userId").(int32) {
		return fiber.ErrForbidden
	}

	group.Name = req.Name

	repository.UpdateGroup(group)

	var users []response.User
	for _, user := range group.Users {
		users = append(users, response.User{
			Uuid: user.UUID,
			Name: user.Name,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Group{
		Uuid: group.Uuid,
		Name: group.Name,
		Admin: response.User{
			Uuid: group.Admin.UUID,
		},
		Users: users,
	})
}

// DeleteGroup godoc
// @Summary      Delete group
// @Description  delete group by uuid
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        uuid path string true "uuid of the group"
// @Success      204  {object}  nil
// @Failure      403  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{uuid} [delete]
func DeleteGroup(c *fiber.Ctx) error {
	groupUuid := c.Params("group")

	group := repository.FindGroupByUuidAndUserId(groupUuid, c.Locals("userId").(int32))

	if group.AdminId != c.Locals("userId").(int32) {
		return fiber.ErrForbidden
	}

	repository.DeleteGroup(*group)

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

// AddUserToGroup godoc
// @Summary      Add user to group
// @Description  Add user to group by uuid
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        uuid path string true "uuid of the group"
// @Param        req body request.UserToGroup true "body"
// @Success      200  {object}  nil
// @Failure      400  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{uuid}/add/user [post]
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

// RemoveUserFromGroup godoc
// @Summary      Remove user from group
// @Description  Remove user from group by uuid
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        uuid path string true "uuid of the group"
// @Param        userId path string true "uuid of the user"
// @Success      204  {object}  nil
// @Failure      403  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{uuid}/remove/user/{userId} [delete]
func RemoveUserFromGroup(c *fiber.Ctx) error {
	groupUuid := c.Params("group")
	userUuid := c.Params("user")

	group := repository.FindGroupByUuidAndUserId(groupUuid, c.Locals("userId").(int32))
	user := repository.FindUserByUuid(userUuid)

	if group.AdminId != c.Locals("userId").(int32) {
		return fiber.ErrForbidden
	}

	var isRemoved bool
	for i, u := range group.Users {
		if u.UUID == user.UUID {
			group.Users = append(group.Users[:i], group.Users[i+1:]...)
			isRemoved = true

			break
		}
	}

	if !isRemoved {
		return fiber.ErrNotFound
	}

	repository.UpdateGroup(group)

	return c.Status(fiber.StatusNoContent).JSON(nil)
}
