package repositories

import (
    "context"
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
    "github.com/stretchr/testify/mock"
)

type MockArticlesRepository struct {
    mock.Mock
}

func (m *MockArticlesRepository) RollbackTransaction(tx *sqlx.Tx) error {
    args := m.Called(tx)
    return args.Error(1)
}

func (m *MockArticlesRepository) GetArticleBySlug(ctx context.Context, tx *sqlx.Tx, slug string) (*ArticleEntity, error) {
    args := m.Called(ctx, tx, slug)
    return args.Get(0).(*ArticleEntity), args.Error(1)
}

func (m *MockArticlesRepository) CreateArticle(ctx context.Context, tx *sqlx.Tx, authorId uuid.UUID, slug, title, description, body string) (*ArticleEntity, error) {
    args := m.Called(ctx, tx, authorId, slug, title, description, body)
    return args.Get(0).(*ArticleEntity), args.Error(1)
}

func (m *MockArticlesRepository) CreateTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error) {
    args := m.Called(ctx, tx, description)
    return args.Get(0).(*TagEntity), args.Error(1)
}

func (m *MockArticlesRepository) CreateArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error) {
    args := m.Called(ctx, tx, tagId, articleId)
    return args.Get(0).(*ArticleTagEntity), args.Error(1)
}

func (m *MockArticlesRepository) GetTag(ctx context.Context, tx *sqlx.Tx, description string) (*TagEntity, error) {
    args := m.Called(ctx, tx, description)
    return args.Get(0).(*TagEntity), args.Error(1)
}

func (m *MockArticlesRepository) GetArticleTag(ctx context.Context, tx *sqlx.Tx, tagId, articleId uuid.UUID) (*ArticleTagEntity, error) {
    args := m.Called(ctx, tx, tagId, articleId)
    return args.Get(0).(*ArticleTagEntity), args.Error(1)
}

func (m *MockArticlesRepository) BeginTransaction(ctx context.Context) (*sqlx.Tx, error) {
    args := m.Called(ctx)
    return args.Get(0).(*sqlx.Tx), args.Error(1)
}

func (m *MockArticlesRepository) CommitTransaction(tx *sqlx.Tx) error {
    args := m.Called(tx)
    return args.Error(0)
}

func (m *MockArticlesRepository) CreateArticleWithTags(ctx context.Context, tx *sqlx.Tx, slug, title, description, body, string, tagList []string) (*ArticleEntity, error) {
    args := m.Called(ctx, tx, slug, title, description, body, tagList)
    return args.Get(0).(*ArticleEntity), args.Error(1)
}
