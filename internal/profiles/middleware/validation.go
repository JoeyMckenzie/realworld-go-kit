package middleware

import (
    "context"
    "github.com/go-kit/log"
    "github.com/go-playground/validator/v10"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

type articlesServiceValidationMiddleware struct {
    validation *validator.Validate
    next       core.ArticlesService
}

func NewArticlesServiceValidationMiddleware(logger log.Logger) core.ArticlesServiceMiddleware {
    return func(next core.ArticlesService) core.ArticlesService {
        return &articlesServiceValidationMiddleware{
            next: next,
        }
    }
}

func (mw articlesServiceValidationMiddleware) CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error) {
    if err := mw.validation.StructCtx(ctx, request); err != nil {
        return &domain.Article{}, shared.MakeValidationError(err)
    }

    return mw.next.CreateArticle(ctx, request, authorId)
}
