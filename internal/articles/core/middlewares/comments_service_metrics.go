package middlewares

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"time"
)

type commentsServiceMetricsMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	service        core.CommentsService
}

func NewCommentsMetricsMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) core.CommentsServiceMiddleware {
	return func(next core.CommentsService) core.CommentsService {
		return &commentsServiceMetricsMiddleware{
			requestCount:   requestCount,
			requestLatency: requestLatency,
			service:        next,
		}
	}
}

func (mw *commentsServiceMetricsMiddleware) AddComment(ctx context.Context, request *domain.AddArticleCommentServiceRequest) (comment *domain.CommentDto, err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "AddComment", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.AddComment(ctx, request)
}

func (mw *commentsServiceMetricsMiddleware) DeleteComment(ctx context.Context, request *domain.DeleteArticleCommentServiceRequest) (err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "DeleteComment", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.DeleteComment(ctx, request)
}

func (mw *commentsServiceMetricsMiddleware) GetArticleComments(ctx context.Context, request *domain.GetCommentsServiceRequest) (comments []*domain.CommentDto, err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "GetArticleComments", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.GetArticleComments(ctx, request)
}
