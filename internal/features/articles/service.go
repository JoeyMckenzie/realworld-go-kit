package articles

import (
    "context"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "golang.org/x/exp/slog"
)

type (
    ArticlesService interface {
        CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error)
    }

    articlesService struct {
        logger             *slog.Logger
        articlesRepository repositories.ArticlesRepository
        usersRepository    repositories.UsersRepository
    }

    ArticlesServiceMiddleware func(service ArticlesService) ArticlesService
)

func NewArticlesService(logger *slog.Logger, articlesRepository repositories.ArticlesRepository, usersRepository repositories.UsersRepository) ArticlesService {
    return &articlesService{
        logger:             logger,
        articlesRepository: articlesRepository,
        usersRepository:    usersRepository,
    }
}
