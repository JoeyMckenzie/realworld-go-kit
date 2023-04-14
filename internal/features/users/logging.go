package users

import (
	"context"
	"github.com/google/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
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

func (mw *usersServiceLoggingMiddleware) Register(ctx context.Context, request domain.AuthenticationRequest[domain.RegisterUserRequest]) (user *domain.User, err error) {
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

func (mw *usersServiceLoggingMiddleware) Login(ctx context.Context, request domain.AuthenticationRequest[domain.LoginUserRequest]) (user *domain.User, err error) {
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

func (mw *usersServiceLoggingMiddleware) Update(ctx context.Context, request domain.AuthenticationRequest[domain.UpdateUserRequest], id uuid.UUID) (user *domain.User, err error) {
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

func (mw *usersServiceLoggingMiddleware) Get(ctx context.Context, id uuid.UUID) (user *domain.User, err error) {
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