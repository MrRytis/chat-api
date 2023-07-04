package exception

import "net/http"

func NewBadRequest(message string) {
	err := NewException(http.StatusBadRequest, http.StatusBadRequest, message, &[]Error{})

	panic(err)
}

func NewUnauthorized(message string) {
	err := NewException(http.StatusUnauthorized, http.StatusUnauthorized, message, &[]Error{})

	panic(err)
}

func NewForbidden(message string) {
	err := NewException(http.StatusForbidden, http.StatusForbidden, message, &[]Error{})

	panic(err)
}

func NewNotFound(message string) {
	err := NewException(http.StatusNotFound, http.StatusNotFound, message, &[]Error{})

	panic(err)
}

func NewUnprocessableEntity(message string, errors *[]Error) {
	err := NewException(http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, message, errors)

	panic(err)
}

func NewUpgradeRequired(message string) {
	err := NewException(http.StatusUpgradeRequired, http.StatusUpgradeRequired, message, &[]Error{})

	panic(err)
}

func NewInternalServerError() {
	err := NewException(http.StatusInternalServerError, http.StatusInternalServerError, "Internal Server Error", &[]Error{})

	panic(err)
}
