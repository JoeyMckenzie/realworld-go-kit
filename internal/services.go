package internal

import (
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
)

type ServiceContainer struct {
	UsersService users.UsersService
}

// MakeServiceContainer builds the downstream services used throughout the application.
func MakeServiceContainer(logger log.Logger, db *pgxpool.Pool) *ServiceContainer {
	validation := validator.New()

	var usersService users.UsersService
	{
		usersRepository := users.NewRepository(db)
		usersService = users.NewService(usersRepository)
		usersService = users.NewUsersServiceLoggingMiddleware(logger)(usersService)
		usersService = users.NewUsersServiceValidationMiddleware(validation)(usersService)
	}

	return &ServiceContainer{
		UsersService: usersService,
	}
}
