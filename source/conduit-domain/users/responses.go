package users

import "github.com/joeymckenzie/realworld-go-kit/conduit-domain/shared"

type (
	UserResponse struct {
		User shared.UserDto `json:"user"`
	}

	ProfileResponse struct {
		Profile shared.ProfileDto `json:"profile"`
	}
)

func NewUserDto(email, username, bio, image, token string) *shared.UserDto {
	return &shared.UserDto{
		Email:    email,
		Token:    token,
		Username: username,
		Bio:      bio,
		Image:    image,
	}
}

func NewDefaultUserDto(email, username, token string) *shared.UserDto {
	return &shared.UserDto{
		Email:    email,
		Token:    token,
		Username: username,
		Bio:      "",
		Image:    "",
	}
}
