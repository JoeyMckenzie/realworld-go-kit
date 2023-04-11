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
    ErrMock                      = errors.New("stub err")
    ErrEmailEmpty                = errors.New("email cannot be empty")
    ErrEmailInvalid              = errors.New("email is invalid")
    ErrUsernameNil               = errors.New("username must be provided")
    ErrUsernameEmpty             = errors.New("username cannot be empty")
    ErrPasswordEmpty             = errors.New("password cannot be empty")
    ErrInvalidLoginAttempt       = errors.New("email or password is invalid")
    ErrUserNotFound              = errors.New("user was not found")
    ErrArticlesNotFound          = errors.New("no articles were found")
    ErrCommentNotFound           = errors.New("no comments were found")
    ErrUsernameOrEmailTaken      = errors.New("username or email already exists")
    ErrArticleTitleExists        = errors.New("article with that title already exists")
    ErrInvalidSigningMethod      = errors.New("invalid token signing method")
    ErrUnauthorized              = errors.New("unauthorized attempt to access resource")
    ErrUsernameNotProvided       = errors.New("username was not provided on the request")
    ErrInternalServerError       = errors.New("unexpected error occurred")
    ErrInvalidRequestBody        = errors.New("request body is invalid")
    ErrInvalidLimitOrOffsetValue = errors.New("limit or offset value is not a valid integer")
    ErrInvalidTokenHeader        = errors.New("authorization token is malformed")
    ErrCannotFollowSelf          = errors.New("cannot follow or unfollow self")
    ErrNilInput                  = errors.New("cannot pass nil value")
    ErrForbiddenArticleUpdate    = errors.New("forbidden to update this article")
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
            apiErrors[structField] = []string{MakeFriendlyError(validationErr)}
        }
    }

    return &ApiError{
        Code:   http.StatusUnprocessableEntity,
        Errors: apiErrors,
    }
}

// MakeFriendlyError converts the field level validation messages into a user-friendly string.
func MakeFriendlyError(fieldError validator.FieldError) string {
    switch fieldError.Tag() {
    case "required":
        return fmt.Sprintf("%s is required", fieldError.StructField())
    case "email":
        return fmt.Sprintf("%s is invalid", fieldError.Value())
    }

    return fieldError.Error()
}

func MakeGenericError() *ApiError {
    return &ApiError{
        Code: http.StatusInternalServerError,
        Errors: ApiErrorMap{
            "message": []string{
                ErrInternalServerError.Error(),
            },
        },
    }
}
