package validator

import (
	"fmt"

	"github.com/go-playground/validator"
)

type JSONValidator struct {
	validate *validator.Validate
}

func NewJSONValidator() *JSONValidator {
	validate := validator.New()
	return &JSONValidator{
		validate,
	}
}

// Validate validates the given struct as with the rules defined by https://godoc.org/github.com/go-playground/validator
func (j JSONValidator) Validate(data interface{}) error {
	err := j.validate.Struct(data)
	if err != nil {
		return fmt.Errorf("%w", err.(validator.ValidationErrors))
	}
	return nil
}
