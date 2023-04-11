package users

import (
    "context"
    "github.com/go-kit/log"
    "github.com/jackc/pgx/v5"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
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
    // First, retrieve the user to verify the username and email is available
    existingUserId, err := us.repository.GetUser(ctx, *request.User.Username, *request.User.Email)

    if existingUserId != nil {
        return nil, shared.ErrUsernameOrEmailTaken
    } else if err != nil && err != pgx.ErrNoRows {
        return nil, shared.ErrInternalServerError
    }

    _, err = us.repository.CreateUser(ctx, *request.User.Username, *request.User.Email, *request.User.Password)

    if err != nil {
        return nil, err
    }

    return nil, nil
}
