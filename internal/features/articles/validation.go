package articles

import (
    "context"
    "github.com/go-playground/validator/v10"
    "github.com/gofrs/uuid"
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

func (mw articlesServiceValidationMiddleware) UpdateArticle(ctx context.Context, request domain.UpdateArticleRequest, authorId uuid.UUID) (*domain.Article, error) {
    if err := mw.validation.StructCtx(ctx, request); err != nil {
        return &domain.Article{}, shared.MakeValidationError(err)
    }

    return mw.next.UpdateArticle(ctx, request, authorId)
}

func (mw articlesServiceValidationMiddleware) DeleteArticle(ctx context.Context, slug string, authorId uuid.UUID) error {
    return mw.next.DeleteArticle(ctx, slug, authorId)
}

func (mw articlesServiceValidationMiddleware) ListArticles(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) ([]domain.Article, error) {
    return mw.next.ListArticles(ctx, request, userId)
}

func (mw articlesServiceValidationMiddleware) GetFeed(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) ([]domain.Article, error) {
    return mw.next.ListArticles(ctx, request, userId)
}

func (mw articlesServiceValidationMiddleware) GetArticle(ctx context.Context, slug string, userId uuid.UUID) (*domain.Article, error) {
    return mw.next.GetArticle(ctx, slug, userId)
}

func (mw articlesServiceValidationMiddleware) FavoriteArticle(ctx context.Context, slug string, userId uuid.UUID) (*domain.Article, error) {
    return mw.next.FavoriteArticle(ctx, slug, userId)
}

func (mw articlesServiceValidationMiddleware) UnavoriteArticle(ctx context.Context, slug string, userId uuid.UUID) (*domain.Article, error) {
    return mw.next.UnavoriteArticle(ctx, slug, userId)
}

func (mw articlesServiceValidationMiddleware) GetArticleTags(ctx context.Context) ([]string, error) {
    return mw.next.GetArticleTags(ctx)
}
