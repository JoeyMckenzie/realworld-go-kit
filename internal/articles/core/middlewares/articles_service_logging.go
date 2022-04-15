package middlewares

import (
	"context"
	"fmt"
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

func (mw *articlesServiceLoggingMiddleware) GetArticles(ctx context.Context, request *domain.GetArticlesServiceRequest) (articles *[]domain.ArticleDto, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "CreateArticle",
			"request_time", time.Since(begin),
			"articles_found", fmt.Sprint(articles != nil),
			"error", err,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "GetArticles",
		"request", request.ToSafeLoggingStruct(),
	)

	return mw.next.GetArticles(ctx, request)
}

func (mw *articlesServiceLoggingMiddleware) GetArticle(ctx context.Context, request *domain.GetArticleServiceRequest) (article *domain.ArticleDto, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetArticle",
			"request_time", time.Since(begin),
			"article_found", fmt.Sprint(article != nil),
			"error", err,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "GetArticle",
		"slug", request.Slug,
		"user_id", request.UserId,
	)

	return mw.next.GetArticle(ctx, request)
}

func (mw *articlesServiceLoggingMiddleware) CreateArticle(ctx context.Context, request *domain.UpsertArticleServiceRequest) (article *domain.ArticleDto, err error) {
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

func (mw *articlesServiceLoggingMiddleware) GetFeed(ctx context.Context, request *domain.GetArticlesServiceRequest) (articles *[]domain.ArticleDto, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetFeed",
			"request_time", time.Since(begin),
			"articles_found", fmt.Sprint(articles != nil),
			"error", err,
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "GetFeed",
		"request", request.ToSafeLoggingStruct(),
	)

	return mw.next.GetFeed(ctx, request)
}
