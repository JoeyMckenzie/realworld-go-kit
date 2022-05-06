package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
	"net/http"
	"strings"
)

type ConduitErrorMap map[string][]string

type ConduitError struct {
	Code   int             `json:"-"`
	Errors ConduitErrorMap `json:"errors"`
}

func NewGenericError() *ConduitError {
	return &ConduitError{
		Errors: ConduitErrorMap{
			"message": []string{
				utilities.ErrInternalServerError.Error(),
			},
		},
	}
}

func NewApiErrors(code int, errors map[string][]string) *ConduitError {
	return &ConduitError{
		Code:   code,
		Errors: errors,
	}
}

func NewApiErrorWithContext(code int, context string, err error) *ConduitError {
	return &ConduitError{
		Code: code,
		Errors: map[string][]string{
			context: {
				err.Error(),
			},
		},
	}
}

func NewInternalServerErrorWithContext(context string, err error) *ConduitError {
	return &ConduitError{
		Code: http.StatusInternalServerError,
		Errors: map[string][]string{
			context: {
				err.Error(),
			},
		},
	}
}

func (ae ConduitError) Error() string {
	serialized, _ := json.Marshal(ae)
	return string(serialized)
}

func ToFriendlyError(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fieldError.StructField())
	case "email":
		return fmt.Sprintf("%s is invalid", fieldError.Value())
	}

	return fieldError.Error()
}

func NewValidationError(validationErrors error) *ConduitError {
	apiErrors := make(map[string][]string)

	for _, validationErr := range validationErrors.(validator.ValidationErrors) {
		structField := strings.ToLower(validationErr.StructField())
		if _, exists := apiErrors[structField]; !exists {
			apiErrors[structField] = []string{ToFriendlyError(validationErr)}
		}
	}

	return &ConduitError{
		Code:   http.StatusUnprocessableEntity,
		Errors: apiErrors,
	}
}
