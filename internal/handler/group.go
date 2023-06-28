package handler

import (
	"github.com/MrRytis/chat-api/internal/model/request"
	"github.com/MrRytis/chat-api/internal/service/groupService"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateGroup godoc
// @Summary      Create new group
// @Description  creates new group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        req body request.CreateGroup true "create group"
// @Success      201  {object}  response.Group
// @Failure      400  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups [post]
func CreateGroup(c *fiber.Ctx) error {
	req := new(request.CreateGroup)
	utils.ParseBodyAndValidate(c, req)

	group := groupService.CreateGroup(req.Name, c.Locals("userId").(int32))

	return c.Status(fiber.StatusCreated).JSON(groupService.BuildGroupDTO(group))
}

// GetGroupList godoc
// @Summary      Get group list
// @Description  Get paginated group list
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        page query int false "default 1"
// @Param        limit query int false "default 20"
// @Success      200  {object}  response.GroupList
// @Failure      403  {object}  response.Error
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

	groups := groupService.GetUsersGroups(c.Locals("userId").(int32), page, limit)
	total := groupService.GetTotalUserGroupCount(c.Locals("userId").(int32))

	return c.Status(fiber.StatusOK).JSON(
		groupService.BuildGroupListDTO(
			groups,
			total,
			page,
			limit,
		),
	)
}

// GetGroup godoc
// @Summary      Get single group
// @Description  get single group by uuid
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        uuid path string true "uuid of the group"
// @Success      200  {object}  response.Group
// @Success      403  {object}  response.Error
// @Success      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{uuid} [get]
func GetGroup(c *fiber.Ctx) error {
	groupUuid := c.Params("group")

	group := groupService.GetUserGroupByUuid(groupUuid, c.Locals("userId").(int32))

	return c.Status(fiber.StatusOK).JSON(groupService.BuildGroupDTO(group))
}

// UpdateGroup godoc
// @Summary      Update group
// @Description  update group by uuid
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        req body request.CreateGroup true "update group"
// @Param        uuid path string true "uuid of the group"
// @Success      200  {object}  response.Group
// @Failure      400  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{uuid} [put]
func UpdateGroup(c *fiber.Ctx) error {
	req := new(request.CreateGroup)
	utils.ParseBodyAndValidate(c, req)
	groupUuid := c.Params("group")

	group := groupService.UpdateGroup(groupUuid, c.Locals("userId").(int32), req.Name)

	return c.Status(fiber.StatusOK).JSON(groupService.BuildGroupDTO(group))
}

// DeleteGroup godoc
// @Summary      Delete group
// @Description  delete group by uuid
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        uuid path string true "uuid of the group"
// @Success      204  {object}  nil
// @Failure      403  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{uuid} [delete]
func DeleteGroup(c *fiber.Ctx) error {
	groupUuid := c.Params("group")

	groupService.DeleteGroup(groupUuid, c.Locals("userId").(int32))

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

// AddUserToGroup godoc
// @Summary      Add user to group
// @Description  Add user to group by uuid
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        uuid path string true "uuid of the group"
// @Param        req body request.UserToGroup true "body"
// @Success      200  {object}  response.GroupUserAdded
// @Failure      400  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{uuid}/add/user [post]
func AddUserToGroup(c *fiber.Ctx) error {
	req := new(request.UserToGroup)
	utils.ParseBodyAndValidate(c, req)
	groupUuid := c.Params("group")

	group := groupService.AddUserToGroup(groupUuid, c.Locals("userId").(int32), req.Uuid)

	return c.Status(fiber.StatusOK).JSON(groupService.BuildUserAddedDTO(group, req.Uuid))
}

// RemoveUserFromGroup godoc
// @Summary      Remove user from group
// @Description  Remove user from group by uuid
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
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

	groupService.RemoveUserFromGroup(groupUuid, c.Locals("userId").(int32), userUuid)

	return c.Status(fiber.StatusNoContent).JSON(nil)
}
