package middlewares

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"time"
)

type tagsServiceLoggingMiddleware struct {
	logger log.Logger
	next   core.TagsService
}

func NewTagsServiceLoggingMiddleware(logger log.Logger) core.TagsServiceMiddleware {
	return func(next core.TagsService) core.TagsService {
		return &tagsServiceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw *tagsServiceLoggingMiddleware) GetTags(ctx context.Context) (tags []string, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetTags",
			"request_time", time.Since(begin),
			"tags", len(tags),
			"error", err,
		)
	}(time.Now())

	return mw.next.GetTags(ctx)
}
