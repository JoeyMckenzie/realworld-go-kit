package specs

import (
	"context"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
)

var (
	stubCreateArticleRequest = domain.CreateArticleServiceRequest{
		UserId:      1,
		Title:       "stub title",
		Description: "stub description",
		Body:        "stub body",
		TagList:     &[]string{"stub tag"},
	}
	stubCreateArticleRequestWithoutTagList = domain.CreateArticleServiceRequest{
		UserId:      1,
		Title:       "stub title",
		Description: "stub description",
		Body:        "stub body",
	}
)

type articlesServiceTestFixture struct {
	mockRepository *persistence.MockArticlesRepository
	service        core.ArticlesService
	ctx            context.Context
}

func newArticlesServiceTestFixture() *articlesServiceTestFixture {
	mockRepository := new(persistence.MockArticlesRepository)

	return &articlesServiceTestFixture{
		mockRepository: mockRepository,
		service:        core.NewArticlesServices(nil, mockRepository),
		ctx:            context.Background(),
	}
}
