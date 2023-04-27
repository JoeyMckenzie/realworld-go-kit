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
        GetArticle(ctx context.Context, slug string, userId uuid.UUID) (*domain.Article, error)
        GetFeed(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) ([]domain.Article, error)
        ListArticles(ctx context.Context, request domain.ListArticlesRequest, userId uuid.UUID) ([]domain.Article, error)
        CreateArticle(ctx context.Context, request domain.CreateArticleRequest, authorId uuid.UUID) (*domain.Article, error)
    }

    articlesService struct {
        logger             *slog.Logger
        articlesRepository repositories.ArticlesRepository
        usersRepository    repositories.UsersRepository
        tagsRepository     repositories.TagsRepository
    }

    ArticlesServiceMiddleware func(service ArticlesService) ArticlesService
)

func NewArticlesService(
    logger *slog.Logger,
    articlesRepository repositories.ArticlesRepository,
    usersRepository repositories.UsersRepository,
    tagsRepository repositories.TagsRepository) ArticlesService {
    return &articlesService{
        logger:             logger,
        articlesRepository: articlesRepository,
        usersRepository:    usersRepository,
        tagsRepository:     tagsRepository,
    }
}
