package users

import (
	"context"
	"github.com/go-kit/log/level"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
	"net/http"
)

func (us *userService) Login(ctx context.Context, request domain.AuthenticationRequest[domain.LoginUserRequest]) (*domain.User, error) {
	const loggingSpan string = "login"

	level.Info(us.logger).Log(loggingSpan, "attempting to login user, checking for existing user", "email", request.User.Email)
	existingUsers, err := us.repository.SearchUsers(ctx, "", request.User.Email)

	if len(existingUsers) == 0 {
		level.Error(us.logger).Log(loggingSpan, "user was not found", "email", request.User.Email)
		return &domain.User{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
	} else if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while attempting to query for existing users", "err", err)
		return &domain.User{}, shared.MakeApiError(err)
	}

	existingUser := existingUsers[0]
	level.Info(us.logger).Log(loggingSpan, "user found, attempting to verify password", "username", existingUser.Username, "email", existingUser.Email)
	validLoginAttempt := us.securityService.IsValidPassword(existingUser.Password, request.User.Password)

	if !validLoginAttempt {
		level.Error(us.logger).Log(loggingSpan, "unauthorized attempt to login", "username", existingUser.Username, "email", existingUser.Email)
		return &domain.User{}, shared.MakeApiErrorWithStatus(http.StatusBadRequest, shared.ErrInvalidLoginAttempt)
	}

	level.Info(us.logger).Log(loggingSpan, "user successfully verified, generating token", "username", existingUser.Username, "email", existingUser.Email, "user_id", existingUser.ID.String())
	token, err := us.tokenService.GenerateUserToken(existingUser.ID, existingUser.Email)

	if err != nil {
		level.Error(us.logger).Log(loggingSpan, "error while attempting generate user token", "err", err)
		return &domain.User{}, shared.MakeApiError(err)
	}

	level.Info(us.logger).Log(loggingSpan, "token successfully generated", "username", existingUser.Username, "email", existingUser.Email, "user_id", existingUser.ID.String())

	return existingUser.ToUser(token), nil
}
