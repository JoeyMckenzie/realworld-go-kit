package users

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type usersServiceLoggingMiddleware struct {
	logger log.Logger
	next   UsersService
}

func NewUsersServiceLoggingMiddleware(logger log.Logger) UsersServiceMiddleware {
	return func(next UsersService) UsersService {
		return &usersServiceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw *usersServiceLoggingMiddleware) Register(ctx context.Context, request AuthenticationRequest[RegisterUserRequest]) (user *User, err error) {
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

func (mw *usersServiceLoggingMiddleware) Login(ctx context.Context, request AuthenticationRequest[LoginUserRequest]) (user *User, err error) {
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
