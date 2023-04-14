package core

import (
	"context"
	"github.com/go-kit/log/level"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
	"net/http"
)

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
