package middlewares

import (
	"context"
	"github.com/go-kit/log"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/users/core"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/users/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
)

type usersServiceRequestValidationMiddleware struct {
    logger    log.Logger
    validator *validator.Validate
    next      core.UsersService
}

func NewUsersServiceRequestValidationMiddleware(logger log.Logger, validator *validator.Validate) core.UsersServiceMiddleware {
    return func(next core.UsersService) core.UsersService {
        return &usersServiceRequestValidationMiddleware{
            logger:    logger,
            validator: validator,
            next:      next,
        }
    }
}

func (mw *usersServiceRequestValidationMiddleware) RegisterUser(ctx context.Context, request *domain.RegisterUserServiceRequest) (*domain.UserDto, error) {
    if err := mw.validator.Struct(request); err != nil {
        return nil, api.NewValidationError(err)
    }

    return mw.next.RegisterUser(ctx, request)
}

func (mw *usersServiceRequestValidationMiddleware) LoginUser(ctx context.Context, request *domain.LoginUserServiceRequest) (*domain.UserDto, error) {
    if err := mw.validator.Struct(request); err != nil {
        return nil, api.NewValidationError(err)
    }

    return mw.next.LoginUser(ctx, request)
}

func (mw *usersServiceRequestValidationMiddleware) GetUser(ctx context.Context, userId int) (*domain.UserDto, error) {
    return mw.next.GetUser(ctx, userId)
}

func (mw *usersServiceRequestValidationMiddleware) UpdateUser(ctx context.Context, request *domain.UpdateUserServiceRequest) (*domain.UserDto, error) {
    if err := mw.validator.Struct(request); err != nil {
        return nil, api.NewValidationError(err)
    }

    return mw.next.UpdateUser(ctx, request)
}

func (mw *usersServiceRequestValidationMiddleware) GetUserProfile(ctx context.Context, username string, currentUserId int) (*domain.ProfileDto, error) {
    return mw.next.GetUserProfile(ctx, username, currentUserId)
}

func (mw *usersServiceRequestValidationMiddleware) AddUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*domain.ProfileDto, error) {
    return mw.next.AddUserFollow(ctx, followerUserId, followeeUsername)
}

func (mw *usersServiceRequestValidationMiddleware) RemoveUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*domain.ProfileDto, error) {
    return mw.next.RemoveUserFollow(ctx, followerUserId, followeeUsername)
}
