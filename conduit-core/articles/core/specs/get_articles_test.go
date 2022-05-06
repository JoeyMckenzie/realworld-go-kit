package specs

import (
    "github.com/gosimple/slug"
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/articles/domain"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_WhenArticleExists_ReturnsProperlyMappedArticle(t *testing.T) {
    // Arrange
    request := domain.GetArticleServiceRequest{
        UserId: 1,
        Slug:   slug.Make("testUser1 article"),
    }

    // Act
    response, err := fixture.service.GetArticle(fixture.ctx, &request)

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
    assert.Equal(t, response.Slug, request.Slug)
}

func Test_WhenArticleDoesNotExist_ReturnsNotFound(t *testing.T) {
    // Arrange
    request := domain.GetArticleServiceRequest{
        UserId: 1,
        Slug:   slug.Make("testUser1 article that does not exist"),
    }

    // Act
    response, err := fixture.service.GetArticle(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrArticlesNotFound.Error())
}
