package users

import (
	"context"
	"database/sql"

	"github.com/joeymckenzie/realworld-go-kit/app/users/repository"
)

type UsersService interface {
	Register(ctx context.Context, username, email, password string) error
}

func NewService(db *sql.DB) UsersService {
	queries := repository.New(db)
	return &userService{
		queries: queries,
	}
}

type userService struct {
	queries *repository.Queries
}

func (us *userService) Register(ctx context.Context, username, email, password string) error {
	return nil
}
