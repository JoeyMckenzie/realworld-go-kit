package profiles

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/data"
)

func mapToProfile(u *data.UserEntity, following bool) *domain.Profile {
	return &domain.Profile{
		Username:  u.Username,
		Email:     u.Email,
		Image:     u.Image,
		Bio:       u.Bio,
		Following: following,
	}
}