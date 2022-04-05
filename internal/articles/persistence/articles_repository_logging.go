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

func (mw *articlesRepositoryLoggingMiddleware) GetTags(ctx context.Context, tags []string) (*[]TagEntity, error) {
	//TODO implement me
	panic("implement me")
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

func (mw *articlesRepositoryLoggingMiddleware) CreateArticleTag(ctx context.Context, tagId int, articleId int) (articleTag *ArticleTagEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "CreateArticleTag",
			"exec_time", time.Since(begin),
			"created", fmt.Sprint(articleTag != nil),
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

func (mw *articlesRepositoryLoggingMiddleware) GetArticleTags(ctx context.Context, tags []string) (articleTags *[]ArticleTagEntity, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "GetTags",
			"exec_time", time.Since(begin),
			"found", fmt.Sprint(articleTags != nil),
			"error", fmt.Sprint(err != nil),
		)
	}(time.Now())

	level.Info(mw.logger).Log(
		"method", "GetTags",
		"tags", fmt.Sprintf("[%s]", strings.Join(tags, ",")),
	)

	return mw.next.GetArticleTags(ctx, tags)
}

func (mw *articlesRepositoryLoggingMiddleware) CreateTag(ctx context.Context, tag string) (*TagEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (mw *articlesRepositoryLoggingMiddleware) GetTag(ctx context.Context, tag string) (*TagEntity, error) {
	//TODO implement me
	panic("implement me")
}
