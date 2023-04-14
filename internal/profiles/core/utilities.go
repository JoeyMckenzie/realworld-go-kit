package core

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/profiles"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/data"
)

func mapToProfile(u *data.UserEntity, following bool) *profiles.Profile {
	return &profiles.Profile{
		Username:  u.Username,
		Email:     u.Email,
		Image:     u.Image,
		Bio:       u.Bio,
		Following: following,
	}
}
