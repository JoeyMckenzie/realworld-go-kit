package specs

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/joeymckenzie/realworld-go-kit/ent/enttest"
	"github.com/joeymckenzie/realworld-go-kit/internal"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	articlesPersistence "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	"testing"
)

var (
	StubCreateArticleRequest = domain.UpsertArticleServiceRequest{
		UserId:      1,
		Title:       "stub title",
		Description: "stub description",
		Body:        "stub body",
		TagList:     &[]string{"stub tag"},
	}
	StubCreateArticleRequestWithoutTagList = domain.UpsertArticleServiceRequest{
		UserId:      1,
		Title:       "stub title",
		Description: "stub description",
		Body:        "stub body",
	}
)

type articlesServiceTestFixture struct {
	mockArticlesRepository *articlesPersistence.MockArticlesRepository
	service                core.ArticlesService
	ctx                    context.Context
}

func newArticlesServiceTestFixture(t *testing.T) *articlesServiceTestFixture {
	mockArticlesRepository := new(articlesPersistence.MockArticlesRepository)
	ctx := context.Background()
	testClient := enttest.Open(t, dialect.SQLite, "file:realworld_go_kit?mode=memory&cache=shared&_fk=1")
	internal.SeedData(ctx, testClient)

	return &articlesServiceTestFixture{
		mockArticlesRepository: mockArticlesRepository,
		service:                core.NewArticlesServices(nil, testClient),
		ctx:                    context.Background(),
	}
}
