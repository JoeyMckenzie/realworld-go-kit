package users

import (
	"context"
)

type UsersService interface {
	Register(ctx context.Context, username, email, password string) error
}

func NewService(repository UsersRepository) UsersService {
	return &userService{
		repository: repository,
	}
}

type userService struct {
	repository UsersRepository
}

func (us *userService) Register(ctx context.Context, username, email, password string) error {
	return nil
}
