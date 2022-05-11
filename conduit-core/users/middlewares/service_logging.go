package middlewares

import (
    "context"
    "github.com/go-kit/log"
    "github.com/go-kit/log/level"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/users"
    sharedDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/shared"
    usersDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/users"
    "time"
)

type usersServiceLoggingMiddleware struct {
    logger log.Logger
    next   users.UsersService
}

func NewUsersServiceLoggingMiddleware(logger log.Logger) users.UsersServiceMiddleware {
    return func(next users.UsersService) users.UsersService {
        return &usersServiceLoggingMiddleware{
            logger: logger,
            next:   next,
        }
    }
}

func (mw *usersServiceLoggingMiddleware) RegisterUser(ctx context.Context, request *usersDomain.RegisterUserServiceRequest) (user *sharedDomain.UserDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "RegisterUser",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "RegisterUser",
        "user_request", request.ToSafeLoggingStruct(),
    )

    return mw.next.RegisterUser(ctx, request)
}

func (mw *usersServiceLoggingMiddleware) LoginUser(ctx context.Context, request *usersDomain.LoginUserServiceRequest) (user *sharedDomain.UserDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "LoginUser",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "LoginUser",
        "user_request", request.ToSafeLoggingStruct(),
    )

    return mw.next.LoginUser(ctx, request)
}

func (mw *usersServiceLoggingMiddleware) GetUser(ctx context.Context, userId int) (user *sharedDomain.UserDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "GetUser",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    return mw.next.GetUser(ctx, userId)
}

func (mw *usersServiceLoggingMiddleware) UpdateUser(ctx context.Context, request *usersDomain.UpdateUserServiceRequest) (user *sharedDomain.UserDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "UpdateUser",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "UpdateUser",
        "user_request", request.ToSafeLoggingStruct(),
    )

    return mw.next.UpdateUser(ctx, request)
}

func (mw *usersServiceLoggingMiddleware) GetUserProfile(ctx context.Context, username string, currentUserId int) (profile *sharedDomain.ProfileDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "GetUserProfile",
            "username", username,
            "currentUserId", currentUserId,
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "GetUserProfile",
        "username", username,
    )

    return mw.next.GetUserProfile(ctx, username, currentUserId)
}

func (mw *usersServiceLoggingMiddleware) AddUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (profile *sharedDomain.ProfileDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "AddUserFollow",
            "followerUserId", followerUserId,
            "followeeUserId", followeeUsername,
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    return mw.next.AddUserFollow(ctx, followerUserId, followeeUsername)
}

func (mw *usersServiceLoggingMiddleware) RemoveUserFollow(ctx context.Context, followerUserId int, followeeUsername string) (profile *sharedDomain.ProfileDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "RemoveUserFollow",
            "followerUserId", followerUserId,
            "followeeUserId", followeeUsername,
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    return mw.next.RemoveUserFollow(ctx, followerUserId, followeeUsername)
}
