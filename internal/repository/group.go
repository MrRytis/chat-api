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
