package utils

import (
	"errors"

	"github.com/go-ozzo/ozzo-validation/v4"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationErrors(err error) []FieldError {
	if err == nil {
		return nil
	}

	var ve validation.Errors
	ok := errors.As(err, &ve)
	if !ok {
		return []FieldError{{Field: "error", Message: err.Error()}}
	}

	var out []FieldError
	for field, e := range ve {
		out = append(out, FieldError{
			Field:   field,
			Message: e.Error(),
		})
	}

	return out
}
