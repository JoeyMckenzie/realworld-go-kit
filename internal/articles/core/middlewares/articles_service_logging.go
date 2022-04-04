package middlewares

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"time"
)

type articlesServiceLoggingMiddleware struct {
	logger log.Logger
	next   core.ArticlesService
}

func NewArticlesServiceLoggingMiddleware(logger log.Logger) core.ArticlesServiceMiddleware {
	return func(next core.ArticlesService) core.ArticlesService {
		return &articlesServiceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw *articlesServiceLoggingMiddleware) CreateArticle(ctx context.Context, request *domain.CreateArticleServiceRequest) (article *domain.ArticleDto, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "CreateArticle",
			"request_time", time.Since(begin),
			"error", err,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "CreateArticle",
		"request", request.ToSafeLoggingStruct(),
	)

	return mw.next.CreateArticle(ctx, request)
}
