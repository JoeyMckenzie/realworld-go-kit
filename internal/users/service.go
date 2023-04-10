package users

import (
	"context"
	"github.com/go-kit/log"
)

type (
	UsersService interface {
		Register(ctx context.Context, request AuthenticationRequest) (*User, error)
	}

	userService struct {
		logger     log.Logger
		repository UsersRepository
	}

	UsersServiceMiddleware func(service UsersService) UsersService
)

func NewService(logger log.Logger, repository UsersRepository) UsersService {
	return &userService{
		logger:     logger,
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
