package request

type CreateGroup struct {
	Name string `json:"name" validate:"required"`
}

type UserToGroup struct {
	Uuid string `json:"uuid" validate:"required"`
}
