package specs

import (
	"context"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/core"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	articlesPersistence "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	usersPersistence "github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
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
	mockArticlesRepository *articlesPersistence.MockArticlesRepository
	mockUsersRepository    *usersPersistence.MockUsersRepository
	service                core.ArticlesService
	ctx                    context.Context
}

func newArticlesServiceTestFixture() *articlesServiceTestFixture {
	mockArticlesRepository := new(articlesPersistence.MockArticlesRepository)
	mockUsersRepository := new(usersPersistence.MockUsersRepository)

	return &articlesServiceTestFixture{
		mockArticlesRepository: mockArticlesRepository,
		mockUsersRepository: mockUsersRepository,
		service:                core.NewArticlesServices(nil, mockArticlesRepository, mockUsersRepository),
		ctx:                    context.Background(),
	}
}
