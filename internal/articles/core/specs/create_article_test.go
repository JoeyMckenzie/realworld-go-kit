package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	"github.com/joeymckenzie/realworld-go-kit/pkg/api"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_CreateArticle_WhenRequestIsValid_ReturnsSuccessfulResponse(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(persistence.StubArticle, nil)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequestWithoutTagList)

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &domain.ArticleDto{}, response)
	assert.Nil(t, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticle", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetTag", 0)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateTag", 0)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticleTag", 0)
}

func Test_CreateArticle_WhenRequestIsValidWithTagListAndTagExists_CallsCreateArticleTagWithoutCreateTag(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(persistence.StubArticle, nil)

	fixture.mockRepository.
		On("GetTag", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.StubTag, nil)

	fixture.mockRepository.
		On("CreateArticleTag", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return(persistence.StubArticleTag, nil)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &domain.ArticleDto{}, response)
	assert.Nil(t, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticle", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticleTag", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetTag", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}

func Test_CreateArticle_WhenArticleSlugExists_ReturnsApiErrors(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.StubArticle, nil)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticle", 0)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticleTag", 0)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetTag", 0)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}

func Test_CreateArticle_WhenGetArticleSlugFails_ReturnsWithErrors(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualError(t, err, utilities.ErrMock.Error())
	fixture.mockRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticle", 0)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticleTag", 0)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetTag", 0)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}

func Test_CreateArticle_WhenCreateArticleFails_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticle", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticleTag", 0)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetTag", 0)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}

func Test_CreateArticle_WhenCreateArticleTagFails_ReturnsWithApiErrors(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(persistence.StubArticle, nil)

	fixture.mockRepository.
		On("GetTag", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.StubTag, nil)

	fixture.mockRepository.
		On("CreateArticleTag", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return(nil, utilities.ErrMock)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.IsType(t, &api.ApiErrors{}, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticle", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticleTag", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetTag", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}

func Test_CreateArticle_WhenTagExists_DoesNotCreateTag(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(persistence.StubArticle, nil)

	fixture.mockRepository.
		On("GetTag", fixture.ctx, mock.AnythingOfType("string")).
		Return(persistence.StubTag, nil)

	fixture.mockRepository.
		On("CreateArticleTag", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return(persistence.StubArticleTag, nil)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &stubCreateArticleRequest)

	// Assert
	assert.NotNil(t, response)
	assert.IsType(t, &domain.ArticleDto{}, response)
	assert.Nil(t, err)
	fixture.mockRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticle", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateArticleTag", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "GetTag", 1)
	fixture.mockRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}
