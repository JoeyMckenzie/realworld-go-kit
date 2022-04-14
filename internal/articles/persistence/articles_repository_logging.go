package persistence

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"strings"
	"time"
)

type articlesRepositoryLoggingMiddleware struct {
	logger log.Logger
	next   ArticlesRepository
}

func NewArticlesRepositoryLoggingMiddleware(logger log.Logger) ArticlesRepositoryMiddleware {
	return func(next ArticlesRepository) ArticlesRepository {
		return &articlesRepositoryLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw *articlesRepositoryLoggingMiddleware) GetArticles(ctx context.Context, request *domain.GetArticlesServiceRequest) (articles *[]ArticleEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetArticles",
			"exec_time", time.Since(begin),
			"found", fmt.Sprint(articles != nil),
			"error", fmt.Sprint(err != nil),
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "GetArticles",
		"request", request.ToSafeLoggingStruct(),
	)

	return mw.next.GetArticles(ctx, request)
}

func (mw *articlesRepositoryLoggingMiddleware) CreateArticle(ctx context.Context, userId int, title, slug, description, body string) (article *ArticleEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "CreateArticle",
			"exec_time", time.Since(begin),
			"created", fmt.Sprint(article != nil),
			"error", fmt.Sprint(err != nil),
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "CreateArticle",
		"title", title,
		"description", description,
	)

	return mw.next.CreateArticle(ctx, userId, title, slug, description, body)
}

func (mw *articlesRepositoryLoggingMiddleware) FindArticleBySlug(ctx context.Context, slug string) (article *ArticleEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "FindArticleBySlug",
			"exec_time", time.Since(begin),
			"found", fmt.Sprint(article != nil),
			"error", fmt.Sprint(err != nil),
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "FindArticleBySlug",
		"slug", slug,
	)

	return mw.next.FindArticleBySlug(ctx, slug)
}

func (mw *articlesRepositoryLoggingMiddleware) CreateTag(ctx context.Context, tag string) (createdTag *TagEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "CreateTag",
			"exec_time", time.Since(begin),
			"created", fmt.Sprint(createdTag != nil),
			"error", fmt.Sprint(err != nil),
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "CreateTag",
		"tag", tag,
	)

	return mw.next.CreateTag(ctx, tag)
}

func (mw *articlesRepositoryLoggingMiddleware) CreateArticleTag(ctx context.Context, tagId, articleId int) (createdArticleTag *ArticleTagEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "CreateArticleTag",
			"exec_time", time.Since(begin),
			"created", fmt.Sprint(createdArticleTag != nil),
			"error", fmt.Sprint(err != nil),
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "CreateArticleTag",
		"tagId", tagId,
		"articleId", articleId,
	)

	return mw.next.CreateArticleTag(ctx, tagId, articleId)
}

func (mw *articlesRepositoryLoggingMiddleware) GetTags(ctx context.Context, searchTags []string) (tags *[]TagEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetTags",
			"exec_time", time.Since(begin),
			"found", fmt.Sprint(searchTags != nil),
			"error", fmt.Sprint(err != nil),
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "GetTags",
		"tags", fmt.Sprintf("[%s]", strings.Join(searchTags, ",")),
	)

	return mw.next.GetTags(ctx, searchTags)
}

func (mw *articlesRepositoryLoggingMiddleware) GetArticleTags(ctx context.Context, articleId int) (tags *[]string, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetArticleTags",
			"exec_time", time.Since(begin),
			"found", fmt.Sprint(tags != nil),
			"error", fmt.Sprint(err != nil),
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "GetArticleTags",
		"article_id", articleId,
	)

	return mw.next.GetArticleTags(ctx, articleId)
}
