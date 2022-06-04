package users

import "fmt"

type (
	RegisterUserApiRequest struct {
		User *struct {
			Email    string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"user,omitempty" validate:"required"`
	}

	RegisterUserServiceRequest struct {
		Email    string `validate:"required,email"`
		Username string `validate:"required"`
		Password string `validate:"required"`
	}

	LoginUserApiRequest struct {
		User *struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}

	LoginUserServiceRequest struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}

	UpdateUserApiRequest struct {
		User *struct {
			Email    *string `json:"email,omitempty"`
			Username *string `json:"username,omitempty"`
			Password *string `json:"password,omitempty"`
			Image    *string `json:"image,omitempty"`
			Bio      *string `json:"bio,omitempty"`
		} `json:"user"`
	}

	UpdateUserServiceRequest struct {
		UserId   int     `validate:"required"`
		Email    *string `validate:"email,omitempty"`
		Username *string
		Password *string
		Image    *string
		Bio      *string
	}

	GetUserServiceRequest struct {
		UserId int    `validate:"required"`
		Token  string `validate:"required,jwt"`
	}

	GetUserProfileApiRequest struct {
		ProfileUsername string
		CurrentUserId   int
	}

	GetUserProfileServiceRequest struct {
		Username      string `validate:"required"`
		CurrentUserId *int
	}

	UserFollowApiRequest struct {
		ProfileUsername string
		CurrentUserId   int
	}
)

func (request *RegisterUserApiRequest) ToSafeLoggingStruct() string {
	if request == nil || request.User == nil {
		return "<nil>"
	}

	return fmt.Sprintf("username: %s; email: %s", request.User.Username, request.User.Email)
}

func (request *RegisterUserServiceRequest) ToSafeLoggingStruct() string {
	if request == nil {
		return "<nil>"
	}

	return fmt.Sprintf("username: %s; email: %s", request.Username, request.Email)
}

func (request *LoginUserApiRequest) ToSafeLoggingStruct() string {
	if request == nil || request.User == nil {
		return "<nil>"
	}

	return fmt.Sprintf("email: %s", request.User.Email)
}

func (request *LoginUserServiceRequest) ToSafeLoggingStruct() string {
	if request == nil {
		return "<nil>"
	}

	return fmt.Sprintf("email: %s", request.Email)
}

func (request *UpdateUserApiRequest) ToSafeLoggingStruct() string {
	if request == nil || request.User == nil {
		return "<nil>"
	}

	return fmt.Sprintf("email: %s; username: %s; image: %s; bio: %s", *request.User.Email, *request.User.Username, *request.User.Image, *request.User.Bio)
}

func (request *UpdateUserServiceRequest) ToSafeLoggingStruct() string {
	if request == nil {
		return "<nil>"
	}

	return fmt.Sprintf("userId: %d; email: %s; username: %s; image: %s; bio: %s",
		request.UserId,
		valueOrNil(request.Email),
		valueOrNil(request.Username),
		valueOrNil(request.Image),
		valueOrNil(request.Bio),
	)
}

func valueOrNil(value *string) string {
	if value == nil {
		return "<nil>"
	}

	return *value
}
