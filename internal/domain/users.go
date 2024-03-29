package domain

import "github.com/gofrs/uuid"

type (
	User struct {
		ID       uuid.UUID `json:"-"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
		Token    string    `json:"token"`
		Image    string    `json:"image"`
		Bio      string    `json:"bio"`
	}

	LoginUserRequest struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}

	RegisterUserRequest struct {
		Username string `validate:"required"`
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}

	UpdateUserRequest struct {
		Username string `json:"username,omitempty"`
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
		Image    string `json:"image,omitempty"`
		Bio      string `json:"bio,omitempty"`
	}

	AuthenticationRequest[T LoginUserRequest | RegisterUserRequest | UpdateUserRequest] struct {
		User *T `json:"user" validate:"required"`
	}

	AuthenticationResponse struct {
		User *User `json:"user"`
	}
)
