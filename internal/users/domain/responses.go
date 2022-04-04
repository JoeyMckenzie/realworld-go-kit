package domain

type (
	UserDto struct {
		Email    string `json:"email"`
		Token    string `json:"token"`
		Username string `json:"username"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	}

	UserResponse struct {
		User UserDto `json:"user"`
	}

	ProfileDto struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	}

	ProfileResponse struct {
		Profile ProfileDto `json:"profile"`
	}
)

func NewUserDto(email, username, bio, image, token string) *UserDto {
	return &UserDto{
		Email:    email,
		Token:    token,
		Username: username,
		Bio:      bio,
		Image:    image,
	}
}

func NewDefaultUserDto(email, username, token string) *UserDto {
	return &UserDto{
		Email:    email,
		Token:    token,
		Username: username,
		Bio:      "",
		Image:    "",
	}
}
