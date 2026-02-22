package middlewares

import (
	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	validator *validator.Validate
}

func (v *StructValidator) Validate(out any) error {
	return v.validator.Struct(out)
}

func NewValidator() *StructValidator {
	return &StructValidator{validator: validator.New()}
}
