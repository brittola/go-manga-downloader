package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) []string {
	var errors []string

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{err.Error()}
	}

	for _, e := range validationErrors {
		field := e.Field()
		tag := e.Tag()
		param := e.Param()

		var msg string
		switch tag {
		case "required":
			msg = fmt.Sprintf("%s is required", field)
		case "min":
			msg = fmt.Sprintf("%s must be at least %s", field, param)
		case "max":
			msg = fmt.Sprintf("%s must be at most %s", field, param)
		case "oneof":
			msg = fmt.Sprintf("%s must be one of: %s", field, param)
		default:
			msg = fmt.Sprintf("%s failed on %s validation", field, tag)
		}
		errors = append(errors, msg)
	}

	return errors
}
