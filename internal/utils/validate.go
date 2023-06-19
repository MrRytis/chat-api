package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(req interface{}) error {
	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		msg := "Field validation error:"

		for _, err := range err.(validator.ValidationErrors) {
			msg += fmt.Sprintf(" %s is %s;", err.Field(), err.Tag())
		}

		return fmt.Errorf(msg)
	}

	return nil
}
