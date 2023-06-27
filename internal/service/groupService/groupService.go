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
	groups := repository.FindPagedGroupsByUserId(page, limit, userId)

	return groups
}

func GetTotalUserGroupCount(userId int32) int64 {
	return repository.GetTotalUserGroupCount(userId)
}

func GetUserGroupByUuid(uuid string, userId int32) entity.Group {
	group, err := repository.FindGroupByUuidAndUserId(uuid, userId)
	utils.HandleDbError(err)

	return group
}

func UpdateGroup(groupUuid string, userId int32, name string) entity.Group {
	group, err := repository.FindGroupByUuidAndUserId(groupUuid, userId)
	utils.HandleDbError(err)

	if group.AdminId != userId {
		exception.NewForbidden("You are not the admin of this group")
	}

	group.Name = name
	group = repository.UpdateGroup(group)

	return group
}

func DeleteGroup(groupUuid string, userId int32) {
	group, err := repository.FindGroupByUuidAndUserId(groupUuid, userId)
	utils.HandleDbError(err)

	if group.AdminId != userId {
		exception.NewForbidden("You are not the admin of this group")
	}

	repository.DeleteGroup(group)
}

func AddUserToGroup(groupUuid string, userId int32, userToAddUuid string) entity.Group {
	group, err := repository.FindGroupByUuidAndUserId(groupUuid, userId)
	utils.HandleDbError(err)

	user, err := repository.FindUserByUuid(userToAddUuid)
	utils.HandleDbError(err)

	group.Users = append(group.Users, user)
	group = repository.UpdateGroup(group)

	return group
}

func RemoveUserFromGroup(groupUuid string, userId int32, userToRemoveUuid string) entity.Group {
	group, err := repository.FindGroupByUuidAndUserId(groupUuid, userId)
	utils.HandleDbError(err)

	user, err := repository.FindUserByUuid(userToRemoveUuid)
	utils.HandleDbError(err)

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

func buildGroup(name string, userId int32) entity.Group {
	admin, err := repository.FindUserById(userId)
	utils.HandleDbError(err)

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
