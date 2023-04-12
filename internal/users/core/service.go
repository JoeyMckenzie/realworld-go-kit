package core

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/infrastructure"
	"github.com/joeymckenzie/realworld-go-kit/internal/utilities"
	"net/http"
)

type (
	// UsersService orchestrates all user and profile related operations a user may perform on their account.
	UsersService interface {
		Register(ctx context.Context, request users.AuthenticationRequest[users.RegisterUserRequest]) (*users.User, error)
		Login(ctx context.Context, request users.AuthenticationRequest[users.LoginUserRequest]) (*users.User, error)
		Update(ctx context.Context, request users.AuthenticationRequest[users.UpdateUserRequest], id uuid.UUID) (*users.User, error)
		Get(ctx context.Context, id uuid.UUID) (*users.User, error)
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

func (us *userService) Register(ctx context.Context, request users.AuthenticationRequest[users.RegisterUserRequest]) (*users.User, error) {
	const loggingSpan string = "registration"

	level.Info(us.logger).Log(loggingSpan, "attempting to register new user, checking for existing users", "username", request.User.Username, "email", request.User.Email)
	existingUsers, err := us.repository.SearchUsers(ctx, request.User.Username, request.User.Email)

	if len(existingUsers) != 0 {
		// Technically, there could be at least two entries here - only log out the first one
		existingUser := existingUsers[0]
		level.Error(us.logger).Log(loggingSpan, "username or email is taken", "username", existingUser.Username, "email", existingUser.Email)
		return &users.User{}, shared.MakeApiErrorWithStatus(http.StatusConflict, shared.ErrUsernameOrEmailTaken)
	} else if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while attempting to query for existing users", "err", err)
		return &users.User{}, shared.MakeApiError(err)
	}

	level.Info(us.logger).Log(loggingSpan, "no user clashes found, hashing user password", "username", request.User.Username, "email", request.User.Email)
	hashedPassword, err := us.securityService.HashPassword(request.User.Password)

	if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while attempting to hash user password", "err", err, "username", request.User.Username, "email", request.User.Email)
		return &users.User{}, shared.MakeApiError(err)
	}

	level.Info(us.logger).Log(loggingSpan, "password successfully hashed, creating user", "username", request.User.Username, "email", request.User.Email)
	createdUser, err := us.repository.CreateUser(ctx, request.User.Username, request.User.Email, hashedPassword)

	if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while attempting create user", "err", err)
		return &users.User{}, shared.MakeApiError(err)
	}

	level.Info(us.logger).Log(loggingSpan, "user successfully created, generating token", "username", createdUser.Username, "email", createdUser.Email, "user_id", createdUser.ID.String())
	token, err := us.tokenService.GenerateUserToken(createdUser.ID, createdUser.Email)

	if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while attempting generate user token", "err", err)
		return &users.User{}, shared.MakeApiError(err)
	}

	level.Info(us.logger).Log(loggingSpan, "token successfully generated", "username", createdUser.Username, "email", createdUser.Email, "user_id", createdUser.ID.String())

	return createdUser.ToUser(token), nil
}

func (us *userService) Login(ctx context.Context, request users.AuthenticationRequest[users.LoginUserRequest]) (*users.User, error) {
	const loggingSpan string = "login"

	level.Info(us.logger).Log(loggingSpan, "attempting to login user, checking for existing user", "email", request.User.Email)
	existingUsers, err := us.repository.SearchUsers(ctx, "", request.User.Email)

	if len(existingUsers) == 0 {
		level.Error(us.logger).Log(loggingSpan, "user was not found", "email", request.User.Email)
		return &users.User{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
	} else if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while attempting to query for existing users", "err", err)
		return &users.User{}, shared.MakeApiError(err)
	}

	existingUser := existingUsers[0]
	level.Info(us.logger).Log(loggingSpan, "user found, attempting to verify password", "username", existingUser.Username, "email", existingUser.Email)
	validLoginAttempt := us.securityService.IsValidPassword(existingUser.Password, request.User.Password)

	if !validLoginAttempt {
		level.Error(us.logger).Log(loggingSpan, "unauthorized attempt to login", "username", existingUser.Username, "email", existingUser.Email)
		return &users.User{}, shared.MakeApiErrorWithStatus(http.StatusBadRequest, shared.ErrInvalidLoginAttempt)
	}

	level.Info(us.logger).Log(loggingSpan, "user successfully verified, generating token", "username", existingUser.Username, "email", existingUser.Email, "user_id", existingUser.ID.String())
	token, err := us.tokenService.GenerateUserToken(existingUser.ID, existingUser.Email)

	if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while attempting generate user token", "err", err)
		return &users.User{}, shared.MakeApiError(err)
	}

	level.Info(us.logger).Log(loggingSpan, "token successfully generated", "username", existingUser.Username, "email", existingUser.Email, "user_id", existingUser.ID.String())

	return existingUser.ToUser(token), nil
}

func (us *userService) Update(ctx context.Context, request users.AuthenticationRequest[users.UpdateUserRequest], id uuid.UUID) (*users.User, error) {
	const loggingSpan string = `update_user`

	// First, verify the user exists before attempting to perform any updates
	level.Info(us.logger).Log(loggingSpan, "attempting to update user, checking for existing user", "email", request.User.Email, "id", id.String())
	existingUser, err := us.repository.GetUser(ctx, id)

	if err != nil && err != sql.ErrNoRows {
		level.Error(us.logger).Log(loggingSpan, "error while attempting check for existing user", "err", err, "email", request.User.Email, "id", id.String())
		return &users.User{}, shared.MakeApiError(err)
	} else if err == sql.ErrNoRows {
		level.Error(us.logger).Log(loggingSpan, "user was not found", "email", request.User.Email, "id", id.String())
		return &users.User{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
	}

	// Next, if an existing username or email exists, invalidate the request
	level.Info(us.logger).Log(loggingSpan, "attempting to verify username and email uniqueness", "email", request.User.Email, "username", request.User.Username, "id", id.String())
	existingUsers, err := us.repository.SearchUsers(ctx, request.User.Username, request.User.Email)

	if err != nil && err != sql.ErrNoRows {
		level.Error(us.logger).Log(loggingSpan, "error attempting to verify username and email uniqueness", "err", err, "email", request.User.Email, "username", request.User.Username, "id", id.String())
		return &users.User{}, shared.MakeApiError(err)
	} else if len(existingUsers) > 0 {
		level.Error(us.logger).Log(loggingSpan, "username or email already exists", "err", err, "email", request.User.Email, "username", request.User.Username, "id", id.String())
		return &users.User{}, shared.MakeApiErrorWithStatus(http.StatusConflict, shared.ErrUsernameOrEmailTaken)
	}

	// Next, re-hash the password if one is found on the request
	if request.User.Password != "" {
		level.Info(us.logger).Log(loggingSpan, "user has updated password, rehashing new password", "email", request.User.Email, "id", id.String())
		updatedHashedPassword, err := us.securityService.HashPassword(existingUser.Password)

		if err != nil {
			level.Error(us.logger).Log(loggingSpan, "error while attempting generated an updated password hash", "err", err, "email", request.User.Email, "id", id.String())
			return &users.User{}, shared.MakeApiError(err)
		}

		existingUser.Password = updatedHashedPassword
	}

	existingUser.Username = shared.GetUpdatedValueIfApplicable(request.User.Username, existingUser.Username)
	existingUser.Email = shared.GetUpdatedValueIfApplicable(request.User.Email, existingUser.Email)
	existingUser.Bio = shared.GetUpdatedValueIfApplicable(request.User.Bio, existingUser.Bio)
	existingUser.Image = shared.GetUpdatedValueIfApplicable(request.User.Image, existingUser.Image)

	level.Info(us.logger).Log(loggingSpan, "attempting to update user in the database", "email", request.User.Email, "id", id.String())
	updatedUser, err := us.repository.UpdateUser(
		ctx,
		id,
		existingUser.Username,
		existingUser.Email,
		existingUser.Bio,
		existingUser.Image,
		existingUser.Password)

	if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while updating user in the database", "err", err, "email", existingUser.Email, "id", id.String())
		return &users.User{}, shared.MakeApiError(err)
	}

	level.Info(us.logger).Log(loggingSpan, "user successfully updated, generating new token", "email", existingUser.Email, "id", id.String())
	token, err := us.tokenService.GenerateUserToken(existingUser.ID, existingUser.Email)

	if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while generating new access token", "err", err, "email", existingUser.Email, "id", id.String())
		return &users.User{}, shared.MakeApiError(err)
	}

	return updatedUser.ToUser(token), nil
}

func (us *userService) Get(ctx context.Context, id uuid.UUID) (*users.User, error) {
	const loggingSpan string = "get_user"

	level.Info(us.logger).Log(loggingSpan, "attempting to get existing user", "email", "id", id)
	existingUser, err := us.repository.GetUser(ctx, id)

	if err != nil && err != sql.ErrNoRows {
		level.Error(us.logger).Log(loggingSpan, "error while attempting check for existing user", "err", err, "id", id)
		return &users.User{}, shared.MakeApiError(err)
	} else if err == sql.ErrNoRows {
		level.Error(us.logger).Log(loggingSpan, "user was not found", "email", "id", id)
		return &users.User{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
	}

	level.Info(us.logger).Log(loggingSpan, "user successfully verified, generating new token", "email", existingUser.Email, "id", id)
	token, err := us.tokenService.GenerateUserToken(existingUser.ID, existingUser.Email)

	if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while generating new access token", "err", err, "email", existingUser.Email, "id", id)
		return &users.User{}, shared.MakeApiError(err)
	}

	return existingUser.ToUser(token), nil
}
