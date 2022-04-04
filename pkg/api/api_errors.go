package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"net/http"
	"strings"
)

type ApiError map[string][]string

type ApiErrors struct {
	Code   int      `json:"code,omitempty"`
	Errors ApiError `json:"errors"`
}

func NewGenericError() *ApiErrors {
	return &ApiErrors{
		Errors: ApiError{
			"message": []string{
				utilities.ErrInternalServerError.Error(),
			},
		},
	}
}

func NewApiErrors(code int, errors map[string][]string) *ApiErrors {
	return &ApiErrors{
		Code:   code,
		Errors: errors,
	}
}

func NewApiErrorWithContext(code int, context string, err error) *ApiErrors {
	return &ApiErrors{
		Code: code,
		Errors: map[string][]string{
			context: {
				err.Error(),
			},
		},
	}
}

func NewInternalServerErrorWithContext(context string, err error) *ApiErrors {
	return &ApiErrors{
		Code: http.StatusInternalServerError,
		Errors: map[string][]string{
			context: {
				err.Error(),
			},
		},
	}
}

func (ae ApiErrors) Error() string {
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

func NewValidationError(validationErrors error) *ApiErrors {
	apiErrors := make(map[string][]string)

	for _, validationErr := range validationErrors.(validator.ValidationErrors) {
		structField := strings.ToLower(validationErr.StructField())
		if field, exists := apiErrors[structField]; exists {
			field = append(field, ToFriendlyError(validationErr))
		} else {
			apiErrors[structField] = []string{ToFriendlyError(validationErr)}
		}
	}

	return &ApiErrors{
		Code:   http.StatusUnsupportedMediaType,
		Errors: apiErrors,
	}
}
