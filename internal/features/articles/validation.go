package articles

import (
    "context"
    "github.com/go-playground/validator/v10"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

type articlesServiceValidationMiddleware struct {
    validation *validator.Validate
    next       ArticlesService
}

func NewArticlesServiceValidationMiddleware(validation *validator.Validate) ArticlesServiceMiddleware {
    return func(next ArticlesService) ArticlesService {
        return &articlesServiceValidationMiddleware{
            next:       next,
            validation: validation,
        }
    }
}

func (mw articlesServiceValidationMiddleware) CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error) {
    if err := mw.validation.StructCtx(ctx, request); err != nil {
        return &domain.Article{}, shared.MakeValidationError(err)
    }

    return mw.next.CreateArticle(ctx, request, authorId)
}
