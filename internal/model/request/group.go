package request

type CreateGroup struct {
	Name string `json:"name" validate:"required"`
}
