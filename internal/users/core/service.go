package core

import (
	"context"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/infrastructure"
	"github.com/joeymckenzie/realworld-go-kit/internal/utilities"
)

type (
	// UsersService orchestrates all user and profile related operations a user may perform on their account.
	UsersService interface {
		Register(ctx context.Context, request users.AuthenticationRequest[users.RegisterUserRequest]) (*users.User, error)
		Login(ctx context.Context, request users.AuthenticationRequest[users.LoginUserRequest]) (*users.User, error)
		Update(ctx context.Context, request users.AuthenticationRequest[users.UpdateUserRequest], id uuid.UUID) (*users.User, error)
		Get(ctx context.Context, id uuid.UUID) (*users.User, error)
		GetProfile(ctx context.Context, username string, followeeId uuid.UUID) (*users.Profile, error)
		Follow(ctx context.Context, username string, followeeId uuid.UUID) (*users.Profile, error)
		Unfollow(ctx context.Context, username string, followeeId uuid.UUID) (*users.Profile, error)
	}

	userService struct {
		logger          log.Logger
		repository      infrastructure.UsersRepository
		tokenService    utilities.TokenService
		securityService utilities.SecurityService
	}

	UsersServiceMiddleware func(service UsersService) UsersService
)

func NewService(logger log.Logger, repository infrastructure.UsersRepository, tokenService utilities.TokenService, securityService utilities.SecurityService) UsersService {
	return &userService{
		logger:          logger,
		repository:      repository,
		tokenService:    tokenService,
		securityService: securityService,
	}
}
