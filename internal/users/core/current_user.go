package core

import (
	"context"
	"database/sql"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
	"net/http"
)

func (us *userService) Get(ctx context.Context, id uuid.UUID) (*users.User, error) {
	const loggingSpan string = "get_user"

	level.Info(us.logger).Log(loggingSpan, "attempting to get existing user", "email", "id", id)
	existingUser, err := us.repository.GetUserById(ctx, id)

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
