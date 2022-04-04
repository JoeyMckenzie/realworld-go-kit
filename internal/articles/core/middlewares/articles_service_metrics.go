package middlewares

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"time"
)

type articlesServiceMetricsMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	service        core.ArticlesService
}

func NewArticlesServiceMetrics(requestCount metrics.Counter, requestLatency metrics.Histogram) core.ArticlesServiceMiddleware {
	return func(next core.ArticlesService) core.ArticlesService {
		return &articlesServiceMetricsMiddleware{
			requestCount:   requestCount,
			requestLatency: requestLatency,
			service:        next,
		}
	}
}

func (mw *articlesServiceMetricsMiddleware) CreateArticle(ctx context.Context, request *domain.CreateArticleServiceRequest) (article *domain.ArticleDto, err error) {
	defer func(begin time.Time) {
		labelValues := []string{"method", "CreateArticle", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(labelValues...).Add(1)
		mw.requestLatency.With(labelValues...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.service.CreateArticle(ctx, request)
}
