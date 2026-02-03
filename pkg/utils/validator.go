package utils

import (
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

type ValidatorErrors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(data interface{}) []ValidatorErrors {
	var errors []ValidatorErrors

	err := Validator.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidatorErrors
			element.Field = err.Field()
			element.Message = getErrorMsg(err)
			errors = append(errors, element)
		}
	}

	return errors
}

func getErrorMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "uuid":
		return "Invalid UUID format"
	default:
		return "Invalid value"
	}
}
