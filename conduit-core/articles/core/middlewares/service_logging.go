package middlewares

import (
    "context"
    "fmt"
    "github.com/go-kit/log"
    "github.com/go-kit/log/level"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/articles/core"
    articlesDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/articles"
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

func (mw *articlesServiceLoggingMiddleware) GetArticles(ctx context.Context, request *articlesDomain.GetArticlesServiceRequest) (articles []*articlesDomain.ArticleDto, err error) {
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

func (mw *articlesServiceLoggingMiddleware) GetArticle(ctx context.Context, request *articlesDomain.GetArticleServiceRequest) (article *articlesDomain.ArticleDto, err error) {
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

func (mw *articlesServiceLoggingMiddleware) CreateArticle(ctx context.Context, request *articlesDomain.CreateArticleServiceRequest) (article *articlesDomain.ArticleDto, err error) {
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

func (mw *articlesServiceLoggingMiddleware) GetFeed(ctx context.Context, request *articlesDomain.GetArticlesServiceRequest) (articles []*articlesDomain.ArticleDto, err error) {
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

func (mw *articlesServiceLoggingMiddleware) UpdateArticle(ctx context.Context, request *articlesDomain.UpdateArticleServiceRequest) (article *articlesDomain.ArticleDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "UpdateArticle",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "UpdateArticle",
        "request", request.ToSafeLoggingStruct(),
    )

    return mw.next.UpdateArticle(ctx, request)
}

func (mw *articlesServiceLoggingMiddleware) DeleteArticle(ctx context.Context, request *articlesDomain.DeleteArticleServiceRequest) (err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "DeleteArticle",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "DeleteArticle",
        "request", request.ToSafeLoggingStruct(),
    )

    return mw.next.DeleteArticle(ctx, request)
}

func (mw *articlesServiceLoggingMiddleware) FavoriteArticle(ctx context.Context, request *articlesDomain.ArticleFavoriteServiceRequest) (article *articlesDomain.ArticleDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "FavoriteArticle",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "FavoriteArticle",
        "request", request.ToSafeLoggingStruct(),
    )

    return mw.next.FavoriteArticle(ctx, request)
}

func (mw *articlesServiceLoggingMiddleware) UnfavoriteArticle(ctx context.Context, request *articlesDomain.ArticleFavoriteServiceRequest) (article *articlesDomain.ArticleDto, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "UnfavoriteArticle",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "UnfavoriteArticle",
        "request", request.ToSafeLoggingStruct(),
    )

    return mw.next.UnfavoriteArticle(ctx, request)
}
