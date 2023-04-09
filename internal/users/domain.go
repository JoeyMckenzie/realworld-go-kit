package users

import "github.com/google/uuid"

type (
	User struct {
		Username string
		Email    string
		Token    string
		Image    string
		Bio      string
	}

	UserEntity struct {
		ID        uuid.UUID
		Username  string
		Email     string
		Password  string
		Image     string
		Bio       string
		CreateAt  string
		UpdatedAt string
	}

	UserRequest struct {
		Username *string `json:"username,omitempty"`
		Email    *string `json:"email"`
		Password *string `json:"password"`
	}

	AuthenticationRequest struct {
		User *UserRequest `json:"user"`
	}

	AuthenticationResponse struct {
		User User `json:"user"`
	}
)

func (u UserEntity) ToUser(token string) *User {
	return &User{
		Username: u.Username,
		Email:    u.Email,
		Token:    token,
		Image:    u.Image,
		Bio:      u.Bio,
	}
}
