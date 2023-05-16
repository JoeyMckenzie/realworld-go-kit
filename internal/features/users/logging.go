package users

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"golang.org/x/exp/slog"
)

type usersServiceLoggingMiddleware struct {
	logger *slog.Logger
	next   UsersService
}

func NewUsersServiceLoggingMiddleware(logger *slog.Logger) UsersServiceMiddleware {
	return func(next UsersService) UsersService {
		return &usersServiceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw *usersServiceLoggingMiddleware) Register(ctx context.Context, request domain.AuthenticationRequest[domain.RegisterUserRequest]) (user *domain.User, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoCtx(ctx,
			"Register",
			"request_time", time.Since(begin),
			"error", err,
			"user_created", user != nil,
		)
	}(time.Now())

	mw.logger.InfoCtx(ctx,
		"Register",
		"request", request,
	)

	return mw.next.Register(ctx, request)
}

func (mw *usersServiceLoggingMiddleware) Login(ctx context.Context, request domain.AuthenticationRequest[domain.LoginUserRequest]) (user *domain.User, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoCtx(ctx,
			"Login",
			"request_time", time.Since(begin),
			"error", err,
			"user_verified", user != nil,
		)
	}(time.Now())

	mw.logger.InfoCtx(ctx,
		"Login",
		"request", request,
	)

	return mw.next.Login(ctx, request)
}

func (mw *usersServiceLoggingMiddleware) Update(ctx context.Context, request domain.AuthenticationRequest[domain.UpdateUserRequest], id uuid.UUID) (user *domain.User, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoCtx(ctx,
			"Update",
			"request_time", time.Since(begin),
			"error", err,
			"user_updated", user != nil,
		)
	}(time.Now())

	mw.logger.InfoCtx(ctx,
		"Update",
		"request", request,
	)

	return mw.next.Update(ctx, request, id)
}

func (mw *usersServiceLoggingMiddleware) Get(ctx context.Context, id uuid.UUID) (user *domain.User, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoCtx(ctx,
			"Get",
			"request_time", time.Since(begin),
			"error", err,
			"user_found", user != nil,
		)
	}(time.Now())

	mw.logger.InfoCtx(ctx,
		"Get",
		"id", id,
	)

	return mw.next.Get(ctx, id)
}
