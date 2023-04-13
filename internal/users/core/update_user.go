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
