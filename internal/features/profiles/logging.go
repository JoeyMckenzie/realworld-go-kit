package profiles

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"time"
)

type profileServiceLoggingMiddleware struct {
	logger log.Logger
	next   ProfilesService
}

func NewProfileServiceLoggingMiddleware(logger log.Logger) ProfileServiceMiddleware {
	return func(next ProfilesService) ProfilesService {
		return &profileServiceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw *profileServiceLoggingMiddleware) Follow(ctx context.Context, username string, followeeId uuid.UUID) (profile *domain.Profile, err error) {
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

func (mw *profileServiceLoggingMiddleware) GetProfile(ctx context.Context, username string, followeeId uuid.UUID) (profile *domain.Profile, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetProfile",
			"request_time", time.Since(begin),
			"error", err,
			"profile_found", profile != nil,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "GetProfile",
		"username", username,
		"followee_id", followeeId,
	)

	return mw.next.GetProfile(ctx, username, followeeId)
}

func (mw *profileServiceLoggingMiddleware) Unfollow(ctx context.Context, username string, followeeId uuid.UUID) (profile *domain.Profile, err error) {
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
