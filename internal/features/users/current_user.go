package users

import (
    "context"
    "database/sql"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

func (us *userService) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
    us.logger.InfoCtx(ctx, "attempting to get existing user", "email", "id", id)
    existingUser, err := us.repository.GetUserById(ctx, id)

    if shared.IsValidSqlErr(err) {
        us.logger.ErrorCtx(ctx, "error while attempting check for existing user", "err", err, "id", id)
        return &domain.User{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        us.logger.ErrorCtx(ctx, "user was not found", "email", "id", id)
        return &domain.User{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
    }

    us.logger.InfoCtx(ctx, "user successfully verified, generating new token", "email", existingUser.Email, "id", id)
    token, err := us.tokenService.GenerateUserToken(existingUser.ID, existingUser.Email)

    if err != nil {
        us.logger.ErrorCtx(ctx, "error while generating new access token", "err", err, "email", existingUser.Email, "id", id)
        return &domain.User{}, shared.MakeApiError(err)
    }

    return existingUser.ToUser(token), nil
}
