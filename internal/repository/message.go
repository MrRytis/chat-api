package repository

import (
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/utils"
)

func CreateMessage(message entity.Message) entity.Message {
	utils.Db.Create(&message)

	return message
}

func FindPagedMessagesByGroupUuid(page int, limit int, groupId uint) ([]entity.Message, error) {
	var messages []entity.Message
	err := utils.Db.Model(&entity.Message{}).
		Where("group_id = ?", groupId).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&messages).
		Error

	if err != nil {
		return messages, err
	}

	return messages, nil
}

func GetTotalGroupMessageCount(groupId uint) (int64, error) {
	var total int64
	err := utils.Db.Model(&entity.Message{}).
		Where("group_id = ?", groupId).
		Count(&total).
		Error

	if err != nil {
		return total, err
	}

	return total, nil
}

func FindMessageByUuidAndGroupId(messageUuid string, groupId uint) (entity.Message, error) {
	var message entity.Message

	err := utils.Db.
		Where("uuid = ?", messageUuid).
		Where("group_id = ?", groupId).
		First(&message).
		Error

	if err != nil {
		return message, err
	}

	return message, nil
}

func UpdateMessage(message entity.Message) entity.Message {
	utils.Db.Save(&message)

	return message
}

func DeleteMessage(message entity.Message) {
	utils.Db.Delete(&message)
}
