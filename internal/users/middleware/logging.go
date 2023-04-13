package middleware

import (
	"context"
	"github.com/google/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/users"
	"github.com/joeymckenzie/realworld-go-kit/internal/users/core"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type usersServiceLoggingMiddleware struct {
	logger log.Logger
	next   core.UsersService
}

func NewUsersServiceLoggingMiddleware(logger log.Logger) core.UsersServiceMiddleware {
	return func(next core.UsersService) core.UsersService {
		return &usersServiceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw *usersServiceLoggingMiddleware) Register(ctx context.Context, request users.AuthenticationRequest[users.RegisterUserRequest]) (user *users.User, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Register",
			"request_time", time.Since(begin),
			"error", err,
			"user_created", user != nil,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "Register",
		"request", request,
	)

	return mw.next.Register(ctx, request)
}

func (mw *usersServiceLoggingMiddleware) Login(ctx context.Context, request users.AuthenticationRequest[users.LoginUserRequest]) (user *users.User, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Login",
			"request_time", time.Since(begin),
			"error", err,
			"user_verified", user != nil,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "Login",
		"request", request,
	)

	return mw.next.Login(ctx, request)
}

func (mw *usersServiceLoggingMiddleware) Update(ctx context.Context, request users.AuthenticationRequest[users.UpdateUserRequest], id uuid.UUID) (user *users.User, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Update",
			"request_time", time.Since(begin),
			"error", err,
			"user_updated", user != nil,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "Update",
		"request", request,
	)

	return mw.next.Update(ctx, request, id)
}

func (mw *usersServiceLoggingMiddleware) Get(ctx context.Context, id uuid.UUID) (user *users.User, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Get",
			"request_time", time.Since(begin),
			"error", err,
			"user_found", user != nil,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "Get",
		"id", id,
	)

	return mw.next.Get(ctx, id)
}

func (mw *usersServiceLoggingMiddleware) Follow(ctx context.Context, username string, followeeId uuid.UUID) (profile *users.Profile, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Follow",
			"request_time", time.Since(begin),
			"error", err,
			"profile_found", profile != nil,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "Follow",
		"username", username,
		"followee_id", followeeId,
	)

	return mw.next.Follow(ctx, username, followeeId)
}

func (mw *usersServiceLoggingMiddleware) Unfollow(ctx context.Context, username string, followeeId uuid.UUID) (profile *users.Profile, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Unfollow",
			"request_time", time.Since(begin),
			"error", err,
			"profile_found", profile != nil,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "Unfollow",
		"username", username,
		"followee_id", followeeId,
	)

	return mw.next.Unfollow(ctx, username, followeeId)
}
