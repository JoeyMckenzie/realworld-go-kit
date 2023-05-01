package users

import (
    "context"
    "database/sql"
    "github.com/gofrs/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
    "net/http"
)

func (us *userService) Update(ctx context.Context, request domain.AuthenticationRequest[domain.UpdateUserRequest], id uuid.UUID) (*domain.User, error) {
    // First, verify the user exists before attempting to perform any updates
    us.logger.InfoCtx(ctx, "attempting to update user, checking for existing user", "email", request.User.Email, "id", id.String())
    existingUser, err := us.repository.GetUserById(ctx, id)

    if shared.IsValidSqlErr(err) {
        us.logger.ErrorCtx(ctx, "error while attempting check for existing user", "err", err, "email", request.User.Email, "id", id.String())
        return &domain.User{}, shared.MakeApiError(err)
    } else if err == sql.ErrNoRows {
        us.logger.ErrorCtx(ctx, "user was not found", "email", request.User.Email, "id", id.String())
        return &domain.User{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
    }

    usernameRequiresUpdating := request.User.Username != "" && request.User.Username != existingUser.Username
    emailRequiresUpdating := request.User.Email != "" && request.User.Email != existingUser.Email

    if usernameRequiresUpdating || emailRequiresUpdating {
        // Next, if an existing username or email exists, invalidate the request
        us.logger.InfoCtx(ctx, "attempting to verify username and email uniqueness", "email", request.User.Email, "username", request.User.Username, "id", id.String())
        existingUsers, err := us.repository.SearchUsers(ctx, request.User.Username, request.User.Email)

        if shared.IsValidSqlErr(err) {
            us.logger.ErrorCtx(ctx, "error attempting to verify username and email uniqueness", "err", err, "email", request.User.Email, "username", request.User.Username, "id", id.String())
            return &domain.User{}, shared.MakeApiError(err)
        } else if len(existingUsers) > 0 {
            us.logger.ErrorCtx(ctx, "username or email already exists", "err", err, "email", request.User.Email, "username", request.User.Username, "id", id.String())
            return &domain.User{}, shared.MakeApiErrorWithStatus(http.StatusConflict, shared.ErrUsernameOrEmailTaken)
        }
    }

    // Next, re-hash the password if one is found on the request
    if request.User.Password != "" {
        us.logger.InfoCtx(ctx, "user has updated password, rehashing new password", "email", request.User.Email, "id", id.String())
        updatedHashedPassword, err := us.securityService.HashPassword(existingUser.Password)

        if err != nil {
            us.logger.ErrorCtx(ctx, "error while attempting generated an updated password hash", "err", err, "email", request.User.Email, "id", id.String())
            return &domain.User{}, shared.MakeApiError(err)
        }

        existingUser.Password = updatedHashedPassword
    }

    existingUser.Username = shared.GetUpdatedValueIfApplicable(request.User.Username, existingUser.Username)
    existingUser.Email = shared.GetUpdatedValueIfApplicable(request.User.Email, existingUser.Email)
    existingUser.Bio = shared.GetUpdatedValueIfApplicable(request.User.Bio, existingUser.Bio)
    existingUser.Image = shared.GetUpdatedValueIfApplicable(request.User.Image, existingUser.Image)

    us.logger.InfoCtx(ctx, "attempting to update user in the data", "email", request.User.Email, "id", id.String())
    updatedUser, err := us.repository.UpdateUser(
        ctx,
        id,
        existingUser.Username,
        existingUser.Email,
        existingUser.Bio,
        existingUser.Image,
        existingUser.Password)

    if err != nil {
        us.logger.ErrorCtx(ctx, "error while updating user in the data", "err", err, "email", existingUser.Email, "id", id.String())
        return &domain.User{}, shared.MakeApiError(err)
    }

    us.logger.InfoCtx(ctx, "user successfully updated, generating new token", "email", existingUser.Email, "id", id.String())
    token, err := us.tokenService.GenerateUserToken(existingUser.ID, existingUser.Email)

    if err != nil {
        us.logger.ErrorCtx(ctx, "error while generating new access token", "err", err, "email", existingUser.Email, "id", id.String())
        return &domain.User{}, shared.MakeApiError(err)
    }

    return updatedUser.ToUser(token), nil
}
