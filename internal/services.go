package internal

import (
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/infrastructure"
)

type ServiceContainer struct {
	UsersService core.UsersService
}

// MakeServiceContainer builds the downstream services used throughout the application.
func MakeServiceContainer(logger log.Logger, db *sqlx.DB) *ServiceContainer {
	validation := validator.New()

	var usersService core.UsersService
	{
		usersRepository := infrastructure.NewRepository(db)
		tokenService := infrastructure.NewTokenService()
		securityService := infrastructure.NewSecurityService()
		usersService = core.NewService(logger, usersRepository, tokenService, securityService)
		usersService = core.NewUsersServiceLoggingMiddleware(logger)(usersService)
		usersService = core.NewUsersServiceValidationMiddleware(validation)(usersService)
	}

	return &ServiceContainer{
		UsersService: usersService,
	}
}
