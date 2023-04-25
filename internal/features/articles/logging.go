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

func NewProfileServiceLoggingMiddleware(logger *slog.Logger) ArticlesServiceMiddleware {
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
            "method", "CreateArticle",
            "request_time", time.Since(begin),
            "error", err,
            "article_created", article != nil,
        )
    }(time.Now())

    mw.logger.InfoCtx(ctx,
        "method", "CreateArticleWithTags",
        "title", request.Article.Title,
        "description", request.Article.Description,
        "body", request.Article.Body,
        "tag_list", request.Article.TagList,
        "author_id", authorId,
    )

    return mw.next.CreateArticle(ctx, request, authorId)
}
