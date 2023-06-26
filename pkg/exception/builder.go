package exception

func NewException(status int, code int, message string, errors *[]Error) *Exception {
	err := &Exception{
		StatusCode: status,
		Code:       code,
		Message:    message,
		Errors:     *errors,
	}

	return err
}

func NewError(field string, message string) *Error {
	err := &Error{
		Field:   field,
		Message: message,
	}

	return err
}
