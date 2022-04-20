package specs

import (
    "github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
    articlesPersistence "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

var (
    createArticleRequestStub = &domain.UpsertArticleServiceRequest{
        UserId:      1,
        Title:       "stub title",
        Description: "stub description",
        Body:        "stub body",
        TagList:     &[]string{"stub tag"},
    }

    createArticleRequestStubWithoutTagList = &domain.UpsertArticleServiceRequest{
        UserId:      1,
        Title:       "stub title",
        Description: "stub description",
        Body:        "stub body",
    }
)

func TestArticlesService(t *testing.T) {
    t.Parallel()
    t.Run("testFuck", testFuck)
}

func testFuck(t *testing.T) {

}

func Test_CreateArticle_GivenValidRequestWithoutExistingTags_ReturnsSuccessfulResponseAndInsertsNewTags(t *testing.T) {
    // Arrange
    fixture := newArticlesServiceTestFixture(t)

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
        On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
        Return(articlesPersistence.StubArticle, nil)

    fixture.mockArticlesRepository.
        On("CreateArticleTag", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
        Return(articlesPersistence.StubArticleTag, nil)

    // Act
    response, err := fixture.service.CreateArticle(fixture.ctx, createArticleRequestStub)

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
    fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
    fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTags", 1)
    fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 1)
    fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticleTag", 1)
}

func Test_CreateArticle_GivenValidRequestWithoutTags_ReturnsSuccessfulResponseWithoutCreatingTag(t *testing.T) {
    // Arrange
    fixture := newArticlesServiceTestFixture(t)

    fixture.mockArticlesRepository.
        On("FindArticleBySlug", fixture.ctx, mock.AnythingOfType("string")).
        Return(nil, nil)

    fixture.mockArticlesRepository.
        On("GetTags", fixture.ctx, mock.AnythingOfType("[]string")).
        Return(nil, nil)

    fixture.mockArticlesRepository.
        On("CreateArticle", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
        Return(articlesPersistence.StubArticle, nil)

    fixture.mockArticlesRepository.
        On("CreateArticleTag", fixture.ctx, mock.AnythingOfType("int"), mock.AnythingOfType("int")).
        Return(articlesPersistence.StubArticleTag, nil)

    // Act
    response, err := fixture.service.CreateArticle(fixture.ctx, createArticleRequestStubWithoutTagList)

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
    fixture.mockArticlesRepository.AssertNumberOfCalls(t, "FindArticleBySlug", 1)
    fixture.mockArticlesRepository.AssertNumberOfCalls(t, "GetTags", 0)
    fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateTag", 0)
    fixture.mockArticlesRepository.AssertNumberOfCalls(t, "CreateArticleTag", 0)
}
