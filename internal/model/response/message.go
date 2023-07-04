package response

import "time"

type Message struct {
	Uuid      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"message"`
	IsEdited  bool      `json:"is_edited"`
	User      User      `json:"user"`
}

type MessageList struct {
	PageNumber int32     `json:"pageNumber"`
	PageSize   int32     `json:"pageSize"`
	ItemsCount int32     `json:"itemsCount"`
	Items      []Message `json:"items"`
}

type GroupMessage struct {
	Message
	Group SimpleGroup `json:"group"`
}
