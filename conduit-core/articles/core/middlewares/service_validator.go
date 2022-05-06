package middlewares

import (
    "context"
    "github.com/go-kit/log"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/articles/core"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/articles/domain"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/api"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
)

type articlesServiceRequestValidationMiddleware struct {
    logger    log.Logger
    validator *validator.Validate
    next      core.ArticlesService
}

func NewArticlesServiceRequestValidationMiddleware(logger log.Logger, validator *validator.Validate) core.ArticlesServiceMiddleware {
    return func(next core.ArticlesService) core.ArticlesService {
        return &articlesServiceRequestValidationMiddleware{
            logger:    logger,
            validator: validator,
            next:      next,
        }
    }
}

func (mw *articlesServiceRequestValidationMiddleware) GetArticles(ctx context.Context, request *domain.GetArticlesServiceRequest) ([]*domain.ArticleDto, error) {
    if request == nil {
        return nil, utilities.ErrNilInput
    }

    return mw.next.GetArticles(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) GetArticle(ctx context.Context, request *domain.GetArticleServiceRequest) (*domain.ArticleDto, error) {
    if request == nil {
        return nil, utilities.ErrNilInput
    }

    return mw.next.GetArticle(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) CreateArticle(ctx context.Context, request *domain.CreateArticleServiceRequest) (*domain.ArticleDto, error) {
    if err := mw.validator.Struct(request); err != nil {
        return nil, api.NewValidationError(err)
    }

    return mw.next.CreateArticle(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) GetFeed(ctx context.Context, request *domain.GetArticlesServiceRequest) ([]*domain.ArticleDto, error) {
    if request == nil {
        return nil, utilities.ErrNilInput
    }

    return mw.next.GetFeed(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) UpdateArticle(ctx context.Context, request *domain.UpdateArticleServiceRequest) (*domain.ArticleDto, error) {
    if err := mw.validator.Struct(request); err != nil {
        return nil, api.NewValidationError(err)
    }

    return mw.next.UpdateArticle(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) DeleteArticle(ctx context.Context, request *domain.DeleteArticleServiceRequest) error {
    if err := mw.validator.Struct(request); err != nil {
        return api.NewValidationError(err)
    }

    return mw.next.DeleteArticle(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) FavoriteArticle(ctx context.Context, request *domain.ArticleFavoriteServiceRequest) (*domain.ArticleDto, error) {
    if err := mw.validator.Struct(request); err != nil {
        return nil, api.NewValidationError(err)
    }

    return mw.next.FavoriteArticle(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) UnfavoriteArticle(ctx context.Context, request *domain.ArticleFavoriteServiceRequest) (*domain.ArticleDto, error) {
    if err := mw.validator.Struct(request); err != nil {
        return nil, api.NewValidationError(err)
    }

    return mw.next.UnfavoriteArticle(ctx, request)
}
