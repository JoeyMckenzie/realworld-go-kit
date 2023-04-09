package users

import (
	"context"
)

type (
	UsersService interface {
		Register(ctx context.Context, request AuthenticationRequest) (*User, error)
	}

	userService struct {
		repository UsersRepository
	}

	UsersServiceMiddleware func(service UsersService) UsersService
)

func NewService(repository UsersRepository) UsersService {
	return &userService{
		repository: repository,
	}
}

func (us *userService) Register(ctx context.Context, request AuthenticationRequest) (*User, error) {
	_, err := us.repository.CreateUser(ctx, *request.User.Username, *request.User.Email, *request.User.Password)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
