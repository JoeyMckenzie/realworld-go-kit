package articles

import (
    "context"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "golang.org/x/exp/slog"
    "time"
)

type articlesServiceLoggingMiddleware struct {
    logger *slog.Logger
    next   ArticlesService
}

func NewArticlesServiceLoggingMiddleware(logger *slog.Logger) ArticlesServiceMiddleware {
    return func(next ArticlesService) ArticlesService {
        return &articlesServiceLoggingMiddleware{
            logger: logger,
            next:   next,
        }
    }
}

func (mw articlesServiceLoggingMiddleware) CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (article *domain.Article, err error) {
    defer func(begin time.Time) {
        mw.logger.InfoCtx(ctx,
            "CreateArticle",
            "request_time", time.Since(begin),
            "error", err,
            "article_created", article != nil,
        )
    }(time.Now())

    mw.logger.InfoCtx(ctx,
        "CreateArticle",
        "title", request.Article.Title,
        "description", request.Article.Description,
        "body", request.Article.Body,
        "tag_list", request.Article.TagList,
        "author_id", authorId,
    )

    return mw.next.CreateArticle(ctx, request, authorId)
}

func (mw articlesServiceLoggingMiddleware) UpdateArticle(ctx context.Context, request domain.UpdateArticleRequest, authorId uuid.UUID) (article *domain.Article, err error) {
    defer func(begin time.Time) {
        mw.logger.InfoCtx(ctx,
            "UpdateArticle",
            "request_time", time.Since(begin),
            "error", err,
            "article_updated", article != nil,
        )
    }(time.Now())

    mw.logger.InfoCtx(ctx,
        "UpdateArticle",
        "title", request.Article.Title,
        "description", request.Article.Description,
        "body", request.Article.Body,
        "author_id", authorId,
    )

    return mw.next.UpdateArticle(ctx, request, authorId)
}

func (mw articlesServiceLoggingMiddleware) DeleteArticle(ctx context.Context, slug string, authorId uuid.UUID) (err error) {
    defer func(begin time.Time) {
        mw.logger.InfoCtx(ctx,
            "DeleteArticle",
            "request_time", time.Since(begin),
            "error", err,
        )
    }(time.Now())

    mw.logger.InfoCtx(ctx,
        "DeleteArticle",
        "author_id", authorId,
        "slug", slug,
    )

    return mw.next.DeleteArticle(ctx, slug, authorId)
}

func (mw articlesServiceLoggingMiddleware) ListArticles(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) (articles []domain.Article, err error) {
    defer func(begin time.Time) {
        mw.logger.InfoCtx(ctx,
            "ListArticles",
            "request_time", time.Since(begin),
            "error", err,
            "articles_found", len(articles),
        )
    }(time.Now())

    mw.logger.InfoCtx(ctx,
        "ListArticles",
        "tag", request.Tag,
        "author", request.Author,
        "favorited", request.Favorited,
        "limit", request.Limit,
        "offset", request.Offset,
        "user_id", userId,
    )

    return mw.next.ListArticles(ctx, request, userId)
}

func (mw articlesServiceLoggingMiddleware) GetFeed(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) (articles []domain.Article, err error) {
    defer func(begin time.Time) {
        mw.logger.InfoCtx(ctx,
            "GetFeed",
            "request_time", time.Since(begin),
            "error", err,
            "articles_found", len(articles),
        )
    }(time.Now())

    mw.logger.InfoCtx(ctx,
        "GetFeed",
        "limit", request.Limit,
        "offset", request.Offset,
        "user_id", userId,
    )

    return mw.next.GetFeed(ctx, request, userId)
}

func (mw articlesServiceLoggingMiddleware) GetArticle(ctx context.Context, slug string, userId uuid.UUID) (article *domain.Article, err error) {
    defer func(begin time.Time) {
        mw.logger.InfoCtx(ctx,
            "GetArticle",
            "request_time", time.Since(begin),
            "error", err,
            "article_found", article != nil,
        )
    }(time.Now())

    mw.logger.InfoCtx(ctx,
        "GetArticle",
        "slug", slug,
        "user_id", userId,
    )

    return mw.next.GetArticle(ctx, slug, userId)
}

func (mw articlesServiceLoggingMiddleware) FavoriteArticle(ctx context.Context, slug string, userId uuid.UUID) (article *domain.Article, err error) {
    defer func(begin time.Time) {
        mw.logger.InfoCtx(ctx,
            "FavoriteArticle",
            "request_time", time.Since(begin),
            "error", err,
            "article_found", article != nil,
        )
    }(time.Now())

    mw.logger.InfoCtx(ctx,
        "FavoriteArticle",
        "slug", slug,
        "user_id", userId,
    )

    return mw.next.FavoriteArticle(ctx, slug, userId)
}

func (mw articlesServiceLoggingMiddleware) UnavoriteArticle(ctx context.Context, slug string, userId uuid.UUID) (article *domain.Article, err error) {
    defer func(begin time.Time) {
        mw.logger.InfoCtx(ctx,
            "UnfavoriteArticle",
            "request_time", time.Since(begin),
            "error", err,
            "article_found", article != nil,
        )
    }(time.Now())

    mw.logger.InfoCtx(ctx,
        "UnfavoriteArticle",
        "slug", slug,
        "user_id", userId,
    )

    return mw.next.UnavoriteArticle(ctx, slug, userId)
}

func (mw articlesServiceLoggingMiddleware) GetArticleTags(ctx context.Context) (tags []string, err error) {
    defer func(begin time.Time) {
        mw.logger.InfoCtx(ctx,
            "GetArticleTags",
            "request_time", time.Since(begin),
            "error", err,
            "article_found", len(tags),
        )
    }(time.Now())

    mw.logger.InfoCtx(ctx, "GetArticleTags")

    return mw.next.GetArticleTags(ctx)
}
