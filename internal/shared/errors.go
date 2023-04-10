package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidRequestBody = errors.New("unable to parse the request body")
)

type (
	// ApiErrorMap represents a map of application errors that can occur while processing requests.
	ApiErrorMap map[string][]string

	// ApiError represents the serialized error response to be propagated back to clients.
	ApiError struct {
		Code   int         `json:"-"`
		Errors ApiErrorMap `json:"errors"`
	}
)

// Error converts an ApiError type into a valid error value.
func (err ApiError) Error() string {
	serialized, _ := json.Marshal(err)
	return string(serialized)
}

// MakeValidationError converts the validator struct errors into a valid API error response type.
func MakeValidationError(validationErrors error) *ApiError {
	apiErrors := make(map[string][]string)

	// Sift through the validation errors and append them to the error response
	for _, validationErr := range validationErrors.(validator.ValidationErrors) {
		structField := strings.ToLower(validationErr.StructField())

		// We'll only append errors to the response map if they have not already been added
		if _, exists := apiErrors[structField]; !exists {
			apiErrors[structField] = []string{ToFriendlyError(validationErr)}
		}
	}

	return &ApiError{
		Code:   http.StatusUnprocessableEntity,
		Errors: apiErrors,
	}
}

// ToFriendlyError converts the field level validation messages into a user-friendly string.
func ToFriendlyError(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fieldError.StructField())
	case "email":
		return fmt.Sprintf("%s is invalid", fieldError.Value())
	}

	return fieldError.Error()
}
