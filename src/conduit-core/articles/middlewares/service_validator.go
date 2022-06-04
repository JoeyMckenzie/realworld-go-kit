package middlewares

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/articles"
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/shared"
	articlesDomain "github.com/joeymckenzie/realworld-go-kit/conduit-domain/articles"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
)

type articlesServiceRequestValidationMiddleware struct {
	logger    log.Logger
	validator *validator.Validate
	next      articles.ArticlesService
}

func NewArticlesServiceRequestValidationMiddleware(logger log.Logger, validator *validator.Validate) articles.ArticlesServiceMiddleware {
	return func(next articles.ArticlesService) articles.ArticlesService {
		return &articlesServiceRequestValidationMiddleware{
			logger:    logger,
			validator: validator,
			next:      next,
		}
	}
}

func (mw *articlesServiceRequestValidationMiddleware) GetArticles(ctx context.Context, request *articlesDomain.GetArticlesServiceRequest) ([]*articlesDomain.ArticleDto, error) {
	if request == nil {
		return nil, utilities.ErrNilInput
	}

	return mw.next.GetArticles(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) GetArticle(ctx context.Context, request *articlesDomain.GetArticleServiceRequest) (*articlesDomain.ArticleDto, error) {
	if request == nil {
		return nil, utilities.ErrNilInput
	}

	return mw.next.GetArticle(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) CreateArticle(ctx context.Context, request *articlesDomain.CreateArticleServiceRequest) (*articlesDomain.ArticleDto, error) {
	if err := mw.validator.Struct(request); err != nil {
		return nil, shared.NewValidationError(err)
	}

	return mw.next.CreateArticle(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) GetFeed(ctx context.Context, request *articlesDomain.GetArticlesServiceRequest) ([]*articlesDomain.ArticleDto, error) {
	if request == nil {
		return nil, utilities.ErrNilInput
	}

	return mw.next.GetFeed(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) UpdateArticle(ctx context.Context, request *articlesDomain.UpdateArticleServiceRequest) (*articlesDomain.ArticleDto, error) {
	if err := mw.validator.Struct(request); err != nil {
		return nil, shared.NewValidationError(err)
	}

	return mw.next.UpdateArticle(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) DeleteArticle(ctx context.Context, request *articlesDomain.DeleteArticleServiceRequest) error {
	if err := mw.validator.Struct(request); err != nil {
		return shared.NewValidationError(err)
	}

	return mw.next.DeleteArticle(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) FavoriteArticle(ctx context.Context, request *articlesDomain.ArticleFavoriteServiceRequest) (*articlesDomain.ArticleDto, error) {
	if err := mw.validator.Struct(request); err != nil {
		return nil, shared.NewValidationError(err)
	}

	return mw.next.FavoriteArticle(ctx, request)
}

func (mw *articlesServiceRequestValidationMiddleware) UnfavoriteArticle(ctx context.Context, request *articlesDomain.ArticleFavoriteServiceRequest) (*articlesDomain.ArticleDto, error) {
	if err := mw.validator.Struct(request); err != nil {
		return nil, shared.NewValidationError(err)
	}

	return mw.next.UnfavoriteArticle(ctx, request)
}
