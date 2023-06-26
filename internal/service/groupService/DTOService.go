package groupService

import (
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/model/response"
	"github.com/MrRytis/chat-api/pkg/exception"
)

func BuildGroupDTO(group entity.Group) response.Group {
	var users []response.User
	for _, user := range group.Users {
		users = append(users, response.User{
			Uuid: user.UUID,
			Name: user.Name,
		})
	}

	return response.Group{
		Uuid: group.Uuid,
		Name: group.Name,
		Admin: response.User{
			Uuid: group.Admin.UUID,
			Name: group.Admin.Name,
		},
		Users: users,
	}
}

func BuildGroupDTOs(groups []entity.Group, total int64, page int, limit int) response.GroupList {
	var groupDTOs []response.Group
	for _, group := range groups {
		groupDTOs = append(groupDTOs, BuildGroupDTO(group))
	}

	return response.GroupList{
		PageNumber: int32(page),
		PageSize:   int32(limit),
		ItemsCount: int32(total),
		Items:      groupDTOs,
	}
}

func BuildUserAddedDTO(group entity.Group, addedUserUuid string) response.GroupUserAdded {
	addedUser := findUserInListByUuid(group, addedUserUuid)

	return response.GroupUserAdded{
		Uuid: group.Uuid,
		User: response.User{
			Uuid: addedUser.UUID,
			Name: addedUser.Name,
		},
		Message: "User added to group",
	}
}

func findUserInListByUuid(group entity.Group, uuid string) entity.User {
	for _, user := range group.Users {
		if user.UUID == uuid {
			return user
		}
	}

	exception.NewInternalServerError()
	return entity.User{}
}
