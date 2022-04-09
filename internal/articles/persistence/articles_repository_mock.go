package persistence

import (
	"context"
	"github.com/stretchr/testify/mock"
	"time"
)

var (
	StubArticle = &ArticleEntity{
		Id:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       "stub title",
		Description: "stub description",
		Body:        "stub body",
	}

	StubArticleTag = &ArticleTagEntity{
		Id:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		TagId:     1,
		ArticleId: StubArticle.Id,
	}

	StubTag = &TagEntity{
		Id:        1,
		CreatedAt: time.Now(),
		Tag:       "stub tag",
	}

	StubAnotherTag = &TagEntity{
		Id:        2,
		CreatedAt: time.Now(),
		Tag:       "another stub tag",
	}
)

type MockArticlesRepository struct {
	mock.Mock
}

func NewMockArticlesRepository() ArticlesRepository {
	return &MockArticlesRepository{}
}

func (m *MockArticlesRepository) GetArticles(ctx context.Context, tag, author, favorited string, limit, offset int) (*[]ArticleEntity, error) {
	args := m.Called(ctx)
	return handleNilArticleMockOrDefault[[]ArticleEntity](args)
}

func (m *MockArticlesRepository) FindArticleBySlug(ctx context.Context, slug string) (*ArticleEntity, error) {
	args := m.Called(ctx, slug)
	return handleNilArticleMockOrDefault[ArticleEntity](args)
}

func (m *MockArticlesRepository) CreateArticle(ctx context.Context, userId int, title, slug, description, body string, tagList []int) (*ArticleEntity, error) {
	args := m.Called(ctx, userId, title, slug, description, body, tagList)
	return handleNilArticleMockOrDefault[ArticleEntity](args)
}

func (m *MockArticlesRepository) GetTags(ctx context.Context, tags []string) (*[]TagEntity, error) {
	args := m.Called(ctx, tags)
	return handleNilArticleMockOrDefault[[]TagEntity](args)
}

func (m *MockArticlesRepository) CreateTag(ctx context.Context, tag string) (*TagEntity, error) {
	args := m.Called(ctx, tag)
	return handleNilArticleMockOrDefault[TagEntity](args)
}

func handleNilArticleMockOrDefault[T ArticleEntity | []ArticleEntity | ArticleTagEntity | []ArticleTagEntity | TagEntity | []TagEntity](args mock.Arguments) (*T, error) {
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*T), args.Error(1)
}
