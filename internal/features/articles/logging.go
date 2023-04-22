package articles

import (
    "context"
    "github.com/go-kit/log"
    "github.com/go-kit/log/level"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "time"
)

type articlesServiceLoggingMiddleware struct {
    logger log.Logger
    next   ArticlesService
}

func NewProfileServiceLoggingMiddleware(logger log.Logger) ArticlesServiceMiddleware {
    return func(next ArticlesService) ArticlesService {
        return &articlesServiceLoggingMiddleware{
            logger: logger,
            next:   next,
        }
    }
}

func (mw articlesServiceLoggingMiddleware) CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (article *domain.Article, err error) {
    defer func(begin time.Time) {
        level.Info(mw.logger).Log(
            "method", "CreateArticle",
            "request_time", time.Since(begin),
            "error", err,
            "article_created", article != nil,
        )
    }(time.Now())

    level.Info(mw.logger).Log(
        "method", "CreateArticleWithTags",
        "title", request.Article.Title,
        "description", request.Article.Description,
        "body", request.Article.Body,
        "tag_list", request.Article.TagList,
        "author_id", authorId,
    )

    return mw.next.CreateArticle(ctx, request, authorId)
}
