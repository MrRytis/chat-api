package exception

func NewBadRequest(message string) {
	err := NewException(400, 400, message, &[]Error{})

	panic(err)
}

func NewUnauthorized(message string) {
	err := NewException(401, 401, message, &[]Error{})

	panic(err)
}

func NewForbidden(message string) {
	err := NewException(403, 403, message, &[]Error{})

	panic(err)
}

func NewNotFound(message string) {
	err := NewException(404, 404, message, &[]Error{})

	panic(err)
}

func NewUnprocessableEntity(message string, errors *[]Error) {
	err := NewException(422, 422, message, errors)

	panic(err)
}

func NewInternalServerError() {
	err := NewException(500, 500, "Internal Server Error", &[]Error{})

	panic(err)
}
