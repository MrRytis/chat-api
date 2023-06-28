package groupService

import (
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/repository"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/MrRytis/chat-api/pkg/exception"
	"github.com/google/uuid"
)

func CreateGroup(name string, userId int32) entity.Group {
	group := buildGroup(name, userId)
	repository.CreateGroup(group)

	return group
}

func GetUsersGroups(userId int32, page int, limit int) []entity.Group {
	groups, err := repository.FindPagedGroupsByUserId(page, limit, userId)
	utils.HandleDbError(err, "Group", "-")

	return groups
}

func GetTotalUserGroupCount(userId int32) int64 {
	total, err := repository.GetTotalUserGroupCount(userId)
	utils.HandleDbError(err, "Group", "-")

	return total
}

func GetUserGroupByUuid(groupUuid string, userId int32) entity.Group {
	group, err := repository.FindGroupByUuid(groupUuid)
	utils.HandleDbError(err, "Group", groupUuid)

	if !IsUserInGroup(group, userId) {
		exception.NewForbidden("You are not in this group")
	}

	return group
}

func UpdateGroup(groupUuid string, userId int32, name string) entity.Group {
	group, err := repository.FindGroupByUuid(groupUuid)
	utils.HandleDbError(err, "Group", groupUuid)

	if !IsUserInGroup(group, userId) {
		exception.NewForbidden("You are not in this group")
	}

	if group.AdminId != userId {
		exception.NewForbidden("You are not the admin of this group")
	}

	group.Name = name
	group = repository.UpdateGroup(group)

	return group
}

func DeleteGroup(groupUuid string, userId int32) {
	group, err := repository.FindGroupByUuid(groupUuid)
	utils.HandleDbError(err, "Group", groupUuid)

	if !IsUserInGroup(group, userId) {
		exception.NewForbidden("You are not in this group")
	}

	if group.AdminId != userId {
		exception.NewForbidden("You are not the admin of this group")
	}

	repository.DeleteGroup(group)
}

func AddUserToGroup(groupUuid string, userId int32, userToAddUuid string) entity.Group {
	group, err := repository.FindGroupByUuid(groupUuid)
	utils.HandleDbError(err, "Group", groupUuid)

	if !IsUserInGroup(group, userId) {
		exception.NewForbidden("You are not in this group")
	}

	user, err := repository.FindUserByUuid(userToAddUuid)
	utils.HandleDbError(err, "User", userToAddUuid)

	group.Users = append(group.Users, user)
	group = repository.UpdateGroup(group)

	return group
}

func RemoveUserFromGroup(groupUuid string, userId int32, userToRemoveUuid string) entity.Group {
	group, err := repository.FindGroupByUuid(groupUuid)
	utils.HandleDbError(err, "Group", groupUuid)

	if !IsUserInGroup(group, userId) {
		exception.NewForbidden("You are not in this group")
	}

	user, err := repository.FindUserByUuid(userToRemoveUuid)
	utils.HandleDbError(err, "User", userToRemoveUuid)

	if group.AdminId != userId {
		exception.NewForbidden("You are not the admin of this group")
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
		exception.NewNotFound("User not found in group")
	}

	group = repository.UpdateGroup(group)

	return group
}

func IsUserInGroup(group entity.Group, userId int32) bool {
	for _, u := range group.Users {
		if u.ID == uint(userId) {
			return true
		}
	}

	return false
}

func buildGroup(name string, userId int32) entity.Group {
	admin, err := repository.FindUserById(userId)
	utils.HandleDbError(err, "User", "-")

	var users []entity.User
	users = append(users, admin)

	return entity.Group{
		Name:    name,
		Uuid:    uuid.New().String(),
		Admin:   admin,
		AdminId: int32(admin.ID),
		Users:   users,
	}
}
