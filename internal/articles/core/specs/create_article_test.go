package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	articlesPersistence "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	usersPersistence "github.com/joeymckenzie/realworld-go-kit/internal/users/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	createArticleRequestStub = &domain.CreateArticleServiceRequest{
		UserId:      1,
		Title:       "stub title",
		Description: "stub description",
		Body:        "stub body",
		TagList:     &[]string{"stub tag"},
	}

	createArticleRequestStubWithoutTagList = &domain.CreateArticleServiceRequest{
		UserId:      1,
		Title:       "stub title",
		Description: "stub description",
		Body:        "stub body",
	}
)

func Test_CreateArticle_GivenValidRequestWithoutExistingTags_ReturnsSuccessfulResponseAndInsertsNewTags(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockUsersRepository.
		On("GetUser", fixture.ctx, mock.AnythingOfType("int")).
		Return(usersPersistence.MockUser, nil)

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockArticlesRepository.
		On("GetTags", fixture.ctx, mock.AnythingOfType("[]string")).
		Return(nil, nil)

	fixture.mockArticlesRepository.
		On("CreateTag", fixture.ctx, mock.AnythingOfType("string")).
		Return(articlesPersistence.StubTag, nil)

	fixture.mockArticlesRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("[]int")).
		Return(articlesPersistence.StubArticle, nil)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, createArticleRequestStub)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	fixture.mockUsersRepository.AssertNumberOfCalls(t, "GetUser", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTags", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 1)
}

func Test_CreateArticle_GivenValidRequestWithoutTags_ReturnsSuccessfulResponseWithoutCreatingTag(t *testing.T) {
	// Arrange
	fixture := newArticlesServiceTestFixture()

	fixture.mockUsersRepository.
		On("GetUser", fixture.ctx, mock.AnythingOfType("int")).
		Return(usersPersistence.MockUser, nil)

	fixture.mockArticlesRepository.
		On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
		Return(nil, nil)

	fixture.mockArticlesRepository.
		On("GetTags", fixture.ctx, mock.AnythingOfType("[]string")).
		Return(nil, nil)

	fixture.mockArticlesRepository.
		On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("[]int")).
		Return(articlesPersistence.StubArticle, nil)

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, createArticleRequestStubWithoutTagList)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
	fixture.mockUsersRepository.AssertNumberOfCalls(t, "GetUser", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTags", 0)
	fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 0)
}
