package middlewares

import (
    "context"
    "fmt"
    "github.com/go-kit/kit/metrics"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/tags/core"
    "time"
)

type tagsServiceMetricsMiddleware struct {
    requestCount   metrics.Counter
    requestLatency metrics.Histogram
    service        core.TagsService
}

func NewTagsServiceMetricsMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) core.TagsServiceMiddleware {
    return func(next core.TagsService) core.TagsService {
        return &tagsServiceMetricsMiddleware{
            requestCount:   requestCount,
            requestLatency: requestLatency,
            service:        next,
        }
    }
}

func (mw *tagsServiceMetricsMiddleware) GetTags(ctx context.Context) (tags []string, err error) {
    defer func(begin time.Time) {
        labelValues := []string{"method", "GetTags", "error", fmt.Sprint(err != nil)}
        mw.requestCount.With(labelValues...).Add(1)
        mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
    }(time.Now())

    return mw.service.GetTags(ctx)
}
