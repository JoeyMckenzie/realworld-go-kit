package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	articlesPersistence "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	usersPersistence "github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_CreateArticle_WhenRequestIsValid_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockUsersRepository.
		On("GetUser", fixture.ctx, mock.AnythingOfType("int")).
		Return(usersPersistence.MockUser, nil)

	fixture.mockArticlesRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("[]int")).
		Return(articlesPersistence.StubArticle, nil)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequestWithoutTagList)

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &domain.ArticleDto{}, response)
	assert.Nil(t, err)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticle", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTag", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticleTag", 0)
}

func Test_CreateArticle_WhenRequestIsValidWithTagListAndTagExists_CallsCreateArticleTagWithoutCreateTag(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockUsersRepository.
		On("GetUser", fixture.ctx, mock.AnythingOfType("int")).
		Return(usersPersistence.MockUser, nil)

	fixture.mockArticlesRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("[]int")).
		Return(articlesPersistence.StubArticle, nil)

	fixture.mockArticlesRepository.
		On("GetTag", fixture.ctx, mock.AnythingOfType("string")).
		Return(articlesPersistence.StubTag, nil)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &domain.ArticleDto{}, response)
	assert.Nil(t, err)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticle", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTags", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}

func Test_CreateArticle_WhenArticleSlugExists_ReturnsApiErrors(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockUsersRepository.
		On("GetUser", fixture.ctx, mock.AnythingOfType("int")).
		Return(usersPersistence.MockUser, nil)

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(articlesPersistence.StubArticle, nil)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticle", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticleTag", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTag", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}

func Test_CreateArticle_WhenGetArticleSlugFails_ReturnsWithErrors(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualError(t, err, utilities.ErrMock.Error())
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticle", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticleTag", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTag", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}

func Test_CreateArticle_WhenCreateArticleFails_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockArticlesRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticle", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticleTag", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTag", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}

func Test_CreateArticle_WhenCreateArticleTagFails_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockArticlesRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(articlesPersistence.StubArticle, nil)

	fixture.mockArticlesRepository.
		On("GetTag", fixture.ctx, mock.AnythingOfType("string")).
		Return(articlesPersistence.StubTag, nil)

	fixture.mockArticlesRepository.
		On("CreateArticleTag", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticle", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticleTag", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTag", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}

func Test_CreateArticle_WhenTagExists_DoesNotCreateTag(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockArticlesRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(articlesPersistence.StubArticle, nil)

	fixture.mockArticlesRepository.
		On("GetTag", fixture.ctx, mock.AnythingOfType("string")).
		Return(articlesPersistence.StubTag, nil)

	fixture.mockArticlesRepository.
		On("CreateArticleTag", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return(articlesPersistence.StubArticleTag, nil)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &domain.ArticleDto{}, response)
	assert.Nil(t, err)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticle", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticleTag", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTag", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}
