package middlewares

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/tags"
	"time"
)

type tagsServiceLoggingMiddleware struct {
    logger log.Logger
    next   tags.TagsService
}

func NewTagsServiceLoggingMiddleware(logger log.Logger) tags.TagsServiceMiddleware {
    return func(next tags.TagsService) tags.TagsService {
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
