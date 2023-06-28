package response

import "time"

type MessageUser struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type Message struct {
	Uuid      string      `json:"uuid"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Content   string      `json:"content"`
	IsEdited  bool        `json:"is_edited"`
	User      MessageUser `json:"user"`
}

type MessageList struct {
	PageNumber int32     `json:"pageNumber"`
	PageSize   int32     `json:"pageSize"`
	ItemsCount int32     `json:"itemsCount"`
	Items      []Message `json:"items"`
}
