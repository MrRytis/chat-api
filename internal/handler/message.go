package handler

import (
	"github.com/MrRytis/chat-api/internal/model/request"
	"github.com/MrRytis/chat-api/internal/service/messageService"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateMessage godoc
// @Summary      Create new message in group
// @Description  creates a new message to specified group
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        req body request.Message true "message"
// @Success      201  {object}  response.Message
// @Failure      400  {object}  response.Error
// @Failure      401  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      422  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{group}/messages [post]
func CreateMessage(c *fiber.Ctx) error {
	req := new(request.Message)
	utils.ParseBodyAndValidate(c, req)

	groupUuid := c.Params("group")
	userId := c.Locals("userId").(int32)

	message := messageService.CreateMessage(userId, groupUuid, req.Content)

	return c.Status(fiber.StatusCreated).JSON(messageService.BuildMessageDTO(message))
}

// GetMessageList godoc
// @Summary      Get list of messages
// @Description  get paginated list of messages
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        group path string true "uuid of the group"
// @Param        page query int false "default 1"
// @Param        limit query int false "default 20"
// @Success      200  {object}  response.Group
// @Failure      401  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{group}/messages [get]
func GetMessageList(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 20)
	page := c.QueryInt("page", 1)
	userId := c.Locals("userId").(int32)
	groupUuid := c.Params("group")

	if limit > 100 {
		limit = 100
	}

	if page < 1 {
		page = 1
	}

	messages := messageService.GetGroupsMessages(userId, groupUuid, page, limit)
	total := messageService.GetTotalGroupMessages(messages[0].GroupId)

	return c.Status(fiber.StatusOK).JSON(
		messageService.BuildMessageListDTO(
			messages,
			total,
			page,
			limit,
		),
	)
}

// GetMessage godoc
// @Summary      Get single message
// @Description  get single message by uuid
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        group path string true "uuid of the group"
// @Param        uuid path string true "uuid of the message"
// @Success      200  {object}  response.Group
// @Failure      401  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{group}/messages/{uuid} [get]
func GetMessage(c *fiber.Ctx) error {
	groupUuid := c.Params("group")
	messageUuid := c.Params("message")
	userId := c.Locals("userId").(int32)

	message := messageService.GetMessageByUuid(userId, groupUuid, messageUuid)

	return c.Status(fiber.StatusOK).JSON(messageService.BuildMessageDTO(message))
}

// UpdateMessage godoc
// @Summary      Update single message
// @Description  update single message by uuid
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        group path string true "uuid of the group"
// @Param        uuid path string true "uuid of the message"
// @Success      200  {object}  response.Group
// @Failure      401  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{group}/messages/{uuid} [put]
func UpdateMessage(c *fiber.Ctx) error {
	req := new(request.Message)
	utils.ParseBodyAndValidate(c, req)

	userId := c.Locals("userId").(int32)
	groupUuid := c.Params("group")
	messageUuid := c.Params("message")

	message := messageService.UpdateMessage(userId, groupUuid, messageUuid, req.Content)

	return c.Status(fiber.StatusOK).JSON(messageService.BuildMessageDTO(message))
}

// DeleteMessage godoc
// @Summary      Delete single message
// @Description  delete single message by uuid
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param		 Authorization header string true "Bearer token"
// @Param        group path string true "uuid of the group"
// @Param        uuid path string true "uuid of the message"
// @Success      200  {object}  response.Group
// @Failure      401  {object}  response.Error
// @Failure      403  {object}  response.Error
// @Failure      404  {object}  response.Error
// @Failure      500  {object}  response.Error
// @Router       /api/v1/groups/{group}/messages/{uuid} [delete]
func DeleteMessage(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int32)
	groupUuid := c.Params("group")
	messageUuid := c.Params("message")

	messageService.DeleteMessage(userId, groupUuid, messageUuid)

	return c.Status(fiber.StatusNoContent).JSON(nil)
}
