package core

import (
    "context"
    "github.com/go-kit/log"
    "github.com/go-kit/log/level"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/articles/data"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/shared"
)

type (
    ArticlesService interface {
        CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error)
    }

    articlesService struct {
        logger     log.Logger
        repository data.ArticlesRepository
    }

    ArticlesServiceMiddleware func(service ArticlesService) ArticlesService
)

func NewArticlesService(logger log.Logger, repository data.ArticlesRepository) ArticlesService {
    return &articlesService{
        logger:     logger,
        repository: repository,
    }
}

func (as *articlesService) CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error) {
    const loggingSpan string = "create_article"

    _, err := as.repository.BeginTransaction(ctx)
    if err != nil {
        level.Error(as.logger).Log(loggingSpan, "error while attempting to begin transaction", "article_title", request.Title, "author_id", authorId)
        return &domain.Article{}, shared.MakeApiError(err)
    }

    return &domain.Article{}, nil
}
