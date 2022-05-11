package middlewares

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/shared"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/users"
	sharedDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/shared"
	usersDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/users"
)

type usersServiceRequestValidationMiddleware struct {
	logger    log.Logger
	validator *validator.Validate
	next      users.UsersService
}

func NewUsersServiceRequestValidationMiddleware(logger log.Logger, validator *validator.Validate) users.UsersServiceMiddleware {
	return func(next users.UsersService) users.UsersService {
		return &usersServiceRequestValidationMiddleware{
			logger:    logger,
			validator: validator,
			next:      next,
		}
	}
}

func (mw *usersServiceRequestValidationMiddleware) RegisterUser(ctx context.Context, request *usersDomain.RegisterUserServiceRequest) (*sharedDomain.UserDto, error) {
	if err := mw.validator.Struct(request); err != nil {
		return nil, shared.NewValidationError(err)
	}

	return mw.next.RegisterUser(ctx, request)
}

func (mw *usersServiceRequestValidationMiddleware) LoginUser(ctx context.Context, request *usersDomain.LoginUserServiceRequest) (*sharedDomain.UserDto, error) {
	if err := mw.validator.Struct(request); err != nil {
		return nil, shared.NewValidationError(err)
	}

	return mw.next.LoginUser(ctx, request)
}

func (mw *usersServiceRequestValidationMiddleware) GetUser(ctx context.Context, userId int) (*sharedDomain.UserDto, error) {
	return mw.next.GetUser(ctx, userId)
}

func (mw *usersServiceRequestValidationMiddleware) UpdateUser(ctx context.Context, request *usersDomain.UpdateUserServiceRequest) (*sharedDomain.UserDto, error) {
	if err := mw.validator.Struct(request); err != nil {
		return nil, shared.NewValidationError(err)
	}

	return mw.next.UpdateUser(ctx, request)
}

func (mw *usersServiceRequestValidationMiddleware) GetUserProfile(ctx context.Context, username string, currentUserId int) (*sharedDomain.ProfileDto, error) {
	return mw.next.GetUserProfile(ctx, username, currentUserId)
}

func (mw *usersServiceRequestValidationMiddleware) AddUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*sharedDomain.ProfileDto, error) {
	return mw.next.AddUserFollow(ctx, followerUserId, followeeUsername)
}

func (mw *usersServiceRequestValidationMiddleware) RemoveUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (*sharedDomain.ProfileDto, error) {
	return mw.next.RemoveUserFollow(ctx, followerUserId, followeeUsername)
}
