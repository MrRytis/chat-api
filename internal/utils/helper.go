package utils

import (
	"github.com/MrRytis/chat-api/pkg/exception"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
)

func ParseBodyAndValidate(c *fiber.Ctx, req interface{}) {
	ParseBody(c, req)
	ValidateRequest(req)
}

func ParseBody(c *fiber.Ctx, req interface{}) {
	if err := c.BodyParser(req); err != nil {
		exception.NewBadRequest("Failed to parse JSON body")
	}
}

func ValidateRequest(req interface{}) {
	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		english := en.New()
		uni := ut.New(english, english)
		trans, _ := uni.GetTranslator("en")
		_ = en_translations.RegisterDefaultTranslations(validate, trans)

		var errors []exception.Error
		for _, err := range err.(validator.ValidationErrors) {
			e := *exception.NewError(err.Field(), err.Translate(trans))

			errors = append(errors, e)
		}

		exception.NewUnprocessableEntity("Validation error", &errors)
	}
}
