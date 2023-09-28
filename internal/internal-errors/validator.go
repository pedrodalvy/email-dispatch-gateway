package internal_errors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()

	err := validate.Struct(s)
	if err == nil {
		return err
	}

	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)

	validationError := validationErrors[0]
	field := strings.ToLower(validationError.Field())

	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "max":
		return errors.New(field + " must be less than " + validationError.Param() + " characters")
	case "min":
		return errors.New(field + " must be more than " + validationError.Param() + " characters")
	case "gte":
		return errors.New(field + " must be greater than or equal to " + validationError.Param())
	default:
		return errors.New(field + " is invalid")
	}
}
