package messageService

import (
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/repository"
	"github.com/MrRytis/chat-api/internal/service/groupService"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/MrRytis/chat-api/pkg/exception"
	"github.com/google/uuid"
)

func CreateMessage(userId int32, groupUuid string, content string) entity.Message {
	group, err := repository.FindGroupByUuid(groupUuid)
	utils.HandleDbError(err, "Group", groupUuid)

	if !groupService.IsUserInGroup(group, userId) {
		exception.NewForbidden("You are not in this group")
	}

	message := entity.Message{
		UUID:     uuid.New().String(),
		UserId:   uint(userId),
		GroupId:  group.ID,
		Content:  content,
		IsEdited: false,
	}

	message = repository.CreateMessage(message)

	return message
}

func GetGroupsMessages(userId int32, groupUuid string, page int, limit int) []entity.Message {
	group, err := repository.FindGroupByUuid(groupUuid)
	utils.HandleDbError(err, "Group", groupUuid)

	if !groupService.IsUserInGroup(group, userId) {
		exception.NewForbidden("You are not in this group")
	}

	messages, err := repository.FindPagedMessagesByGroupUuid(page, limit, group.ID)
	utils.HandleDbError(err, "Message", "-")

	return messages
}

func GetTotalGroupMessages(groupId uint) int64 {
	total, err := repository.GetTotalGroupMessageCount(groupId)
	utils.HandleDbError(err, "Message", "-")

	return total
}

func GetMessageByUuid(userId int32, groupUuid string, messageUuid string) entity.Message {
	group, err := repository.FindGroupByUuid(groupUuid)
	utils.HandleDbError(err, "Group", groupUuid)

	if !groupService.IsUserInGroup(group, userId) {
		exception.NewForbidden("You are not in this group")
	}

	message, err := repository.FindMessageByUuidAndGroupId(messageUuid, group.ID)
	utils.HandleDbError(err, "Message", messageUuid)

	return message
}

func UpdateMessage(userId int32, groupUuid string, messageUuid string, content string) entity.Message {
	group, err := repository.FindGroupByUuid(groupUuid)
	utils.HandleDbError(err, "Group", groupUuid)

	if !groupService.IsUserInGroup(group, userId) {
		exception.NewForbidden("You are not in this group")
	}

	message, err := repository.FindMessageByUuidAndGroupId(messageUuid, group.ID)
	utils.HandleDbError(err, "Message", messageUuid)

	if message.UserId != uint(userId) {
		exception.NewForbidden("You are not the author of this message")
	}

	message.Content = content
	message.IsEdited = true

	message = repository.UpdateMessage(message)

	return message
}

func DeleteMessage(userId int32, groupUuid string, messageUuid string) {
	group, err := repository.FindGroupByUuid(groupUuid)
	utils.HandleDbError(err, "Group", groupUuid)

	if !groupService.IsUserInGroup(group, userId) {
		exception.NewForbidden("You are not in this group")
	}

	message, err := repository.FindMessageByUuidAndGroupId(messageUuid, group.ID)
	utils.HandleDbError(err, "Message", messageUuid)

	if message.UserId != uint(userId) {
		exception.NewForbidden("You are not the author of this message")
	}

	repository.DeleteMessage(message)
}
