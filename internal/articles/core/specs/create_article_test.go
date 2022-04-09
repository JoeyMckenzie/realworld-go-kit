package specs

import (
	articlesPersistence "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	usersPersistence "github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_CreateArticle_GivenValidRequest_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockUsersRepository.
		On("GetUser", fixture.ctx, mock.AnythingOfType("int")).
		Return(usersPersistence.MockUser, nil)

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(articlesPersistence.StubArticle, nil)

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(articlesPersistence.StubArticle, nil)

	// Act

	// Assert
}
