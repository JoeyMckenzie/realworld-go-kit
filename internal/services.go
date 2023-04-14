package internal

import (
    "github.com/go-kit/log"
    "github.com/go-playground/validator/v10"
    "github.com/jmoiron/sqlx"
    profilesCore "github.com/joeymckenzie/realworld-go-kit/internal/profiles/core"
    profilesMiddleware "github.com/joeymckenzie/realworld-go-kit/internal/profiles/middleware"
    usersCore "github.com/joeymckenzie/realworld-go-kit/internal/users/core"
    "github.com/joeymckenzie/realworld-go-kit/internal/users/data"
    usersMiddleware "github.com/joeymckenzie/realworld-go-kit/internal/users/middleware"
    "github.com/joeymckenzie/realworld-go-kit/internal/utilities"
)

type ServiceContainer struct {
    UsersService    usersCore.UsersService
    ProfilesService profilesCore.ProfilesService
}

// NewServiceContainer builds the downstream services used throughout the application.
func NewServiceContainer(logger log.Logger, db *sqlx.DB) *ServiceContainer {
    validation := validator.New()
    usersRepository := data.NewRepository(db)

    var usersService usersCore.UsersService
    {
        tokenService := utilities.NewTokenService()
        securityService := utilities.NewSecurityService()
        usersService = usersCore.NewUsersService(logger, usersRepository, tokenService, securityService)
        usersService = usersMiddleware.NewUsersServiceLoggingMiddleware(logger)(usersService)
        usersService = usersMiddleware.NewUsersServiceValidationMiddleware(validation)(usersService)
    }

    var profilesService profilesCore.ProfilesService
    {
        profilesService = profilesCore.NewProfileService(logger, usersRepository)
        profilesService = profilesMiddleware.NewProfileServiceLoggingMiddleware(logger)(profilesService)
    }

    return &ServiceContainer{
        UsersService:    usersService,
        ProfilesService: profilesService,
    }
}
