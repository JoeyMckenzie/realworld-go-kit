package internal

import (
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
)

type ServiceContainer struct {
	UsersService users.UsersService
}

// MakeServiceContainer builds the downstream services used throughout the application.
func MakeServiceContainer(logger log.Logger, db *sqlx.DB) *ServiceContainer {
	validation := validator.New()

	var usersService users.UsersService
	{
		usersRepository := users.NewRepository(db)
		tokenService := users.NewTokenService()
		securityService := users.NewSecurityService()
		usersService = users.NewService(logger, usersRepository, tokenService, securityService)
		usersService = users.NewUsersServiceLoggingMiddleware(logger)(usersService)
		usersService = users.NewUsersServiceValidationMiddleware(validation)(usersService)
	}

	return &ServiceContainer{
		UsersService: usersService,
	}
}
