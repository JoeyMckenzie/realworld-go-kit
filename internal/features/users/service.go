package users

import (
	"context"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/data"
	utilities2 "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/utilities"
)

type (
	// UsersService orchestrates all user and profile related operations a user may perform on their account.
	UsersService interface {
		Register(ctx context.Context, request domain.AuthenticationRequest[domain.RegisterUserRequest]) (*domain.User, error)
		Login(ctx context.Context, request domain.AuthenticationRequest[domain.LoginUserRequest]) (*domain.User, error)
		Update(ctx context.Context, request domain.AuthenticationRequest[domain.UpdateUserRequest], id uuid.UUID) (*domain.User, error)
		Get(ctx context.Context, id uuid.UUID) (*domain.User, error)
	}

	userService struct {
		logger          log.Logger
		repository      data.UsersRepository
		tokenService    utilities2.TokenService
		securityService utilities2.SecurityService
	}

	UsersServiceMiddleware func(service UsersService) UsersService
)

func NewUsersService(logger log.Logger, repository data.UsersRepository, tokenService utilities2.TokenService, securityService utilities2.SecurityService) UsersService {
	return &userService{
		logger:          logger,
		repository:      repository,
		tokenService:    tokenService,
		securityService: securityService,
	}
}
