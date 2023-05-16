package profiles

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"golang.org/x/exp/slog"
)

type profileServiceLoggingMiddleware struct {
	logger *slog.Logger
	next   ProfilesService
}

func NewProfileServiceLoggingMiddleware(logger *slog.Logger) ProfileServiceMiddleware {
	return func(next ProfilesService) ProfilesService {
		return &profileServiceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw *profileServiceLoggingMiddleware) Follow(ctx context.Context, username string, followeeId uuid.UUID) (profile *domain.Profile, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoCtx(ctx,
			"Follow",
			"request_time", time.Since(begin),
			"error", err,
			"profile_found", profile != nil,
		)
	}(time.Now())

	mw.logger.InfoCtx(ctx,
		"Follow",
		"username", username,
		"followee_id", followeeId,
	)

	return mw.next.Follow(ctx, username, followeeId)
}

func (mw *profileServiceLoggingMiddleware) GetProfile(ctx context.Context, username string, followeeId uuid.UUID) (profile *domain.Profile, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoCtx(ctx,
			"GetProfile",
			"request_time", time.Since(begin),
			"error", err,
			"profile_found", profile != nil,
		)
	}(time.Now())

	mw.logger.InfoCtx(ctx,
		"GetProfile",
		"username", username,
		"followee_id", followeeId,
	)

	return mw.next.GetProfile(ctx, username, followeeId)
}

func (mw *profileServiceLoggingMiddleware) Unfollow(ctx context.Context, username string, followeeId uuid.UUID) (profile *domain.Profile, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoCtx(ctx,
			"Unfollow",
			"request_time", time.Since(begin),
			"error", err,
			"profile_found", profile != nil,
		)
	}(time.Now())

	mw.logger.InfoCtx(ctx,
		"Unfollow",
		"username", username,
		"followee_id", followeeId,
	)

	return mw.next.Unfollow(ctx, username, followeeId)
}
