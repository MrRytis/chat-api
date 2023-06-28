package repository

import (
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/utils"
	"log"
)

func CreateGroup(group entity.Group) entity.Group {
	err := utils.Db.Create(&group).Error
	if err != nil {
		log.Fatal(err, "Error saving group")
	}

	return group
}

func FindPagedGroupsByUserId(page int, limit int, userId int32) ([]entity.Group, error) {
	var groups []entity.Group

	err := utils.Db.
		Joins("JOIN group_users ON group_users.group_id = groups.id").
		Joins("JOIN users ON users.id = group_users.user_id").
		Where("users.id in (?)", userId).
		Preload("Users").
		Preload("Admin").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&groups).
		Error

	if err != nil {
		return groups, err
	}

	return groups, nil
}

func GetTotalUserGroupCount(userId int32) (int64, error) {
	var total int64
	err := utils.Db.
		Model(&entity.Group{}).
		Joins("JOIN group_users ON group_users.group_id = groups.id").
		Joins("JOIN users ON users.id = group_users.user_id").
		Where("users.id in (?)", userId).
		Count(&total).
		Error

	if err != nil {
		return total, err
	}

	return total, nil
}

func FindGroupByUuid(uuid string) (entity.Group, error) {
	var group entity.Group

	err := utils.Db.
		Where("uuid = ?", uuid).
		Preload("Users").
		Preload("Admin").
		First(&group).
		Error

	if err != nil {
		return entity.Group{}, err
	}

	return group, nil
}

func UpdateGroup(group entity.Group) entity.Group {
	err := utils.Db.Save(group).Error
	if err != nil {
		log.Fatal(err, "Error updating group")
	}

	return group
}

func DeleteGroup(group entity.Group) {
	err := utils.Db.Delete(&group).Error
	if err != nil {
		log.Fatal(err, "Error deleting group")
	}
}
