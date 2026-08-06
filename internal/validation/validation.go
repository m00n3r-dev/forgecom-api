package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func Errors(err error) map[string]string {
	errors := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())

		switch err.Tag() {
		case "required":
			errors[field] = "This field is required"

		case "email":
			errors[field] = "Please enter a valid email address"

		case "min":
			errors[field] = "Must be at least " + err.Param() + " characters"

		default:
			errors[field] = "Invalid value"
		}
	}

	return errors
}
