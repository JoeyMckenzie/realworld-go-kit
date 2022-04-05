package persistence

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
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

func (mw *articlesRepositoryLoggingMiddleware) GetArticles(ctx context.Context, tag, author, favorited string, limit, offset int) (articles *[]ArticleEntity, err error) {
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
		"tag", tag,
		"author", author,
		"favorited", favorited,
		"limit", limit,
		"offset", offset,
	)

	return mw.next.GetArticles(ctx, tag, author, favorited, limit, offset)
}

func (mw *articlesRepositoryLoggingMiddleware) CreateArticle(ctx context.Context, userId int, title, slug, description, body string, tagList []int) (article *ArticleEntity, err error) {
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

	return mw.next.CreateArticle(ctx, userId, title, slug, description, body, tagList)
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
