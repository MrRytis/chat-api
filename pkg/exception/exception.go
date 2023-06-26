package exception

import "fmt"

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Exception struct {
	StatusCode int     `json:"-"`
	Code       int     `json:"code"`
	Message    string  `json:"message"`
	Errors     []Error `json:"errors"`
}

func (e *Exception) Error() string {
	return fmt.Sprintf("Error: %s (Code: %d)", e.Message, e.Code)
}
