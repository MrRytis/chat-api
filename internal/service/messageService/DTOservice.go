package messageService

import (
	"github.com/MrRytis/chat-api/internal/entity"
	"github.com/MrRytis/chat-api/internal/model/response"
)

func BuildMessageDTO(message entity.Message) response.Message {
	user := response.MessageUser{
		Uuid: message.User.UUID,
		Name: message.User.Name,
	}

	return response.Message{
		Uuid:      message.UUID,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
		Content:   message.Content,
		IsEdited:  message.IsEdited,
		User:      user,
	}
}

func BuildMessageListDTO(messages []entity.Message, total int64, page int, limit int) response.MessageList {
	var messageDTOs []response.Message
	for _, message := range messages {
		messageDTOs = append(messageDTOs, BuildMessageDTO(message))
	}

	return response.MessageList{
		PageNumber: int32(page),
		PageSize:   int32(limit),
		ItemsCount: int32(total),
		Items:      messageDTOs,
	}
}
