package middlewares

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
)

type articlesServiceRequestValidationMiddleware struct {
	logger    log.Logger
	validator *validator.Validate
	next      core.ArticlesService
}

func (mw *articlesServiceRequestValidationMiddleware) GetArticles(ctx context.Context, request *domain.GetArticlesServiceRequest) ([]*domain.ArticleDto, error) {
	if request == nil {
		return nil, utilities.ErrNilInput
	}

	return mw.next.GetArticles(ctx, request)
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

func (mw *articlesServiceRequestValidationMiddleware) CreateArticle(ctx context.Context, request *domain.UpsertArticleServiceRequest) (*domain.ArticleDto, error) {
	if err := mw.validator.Struct(request); err != nil {
		return nil, api.NewValidationError(err)
	}

	return mw.next.CreateArticle(ctx, request)
}