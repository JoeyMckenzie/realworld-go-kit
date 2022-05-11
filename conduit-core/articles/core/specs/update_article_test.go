package specs

import (
    "github.com/gosimple/slug"
    "github.com/joeymckenzie/realworld-go-kit/conduit-domain/articles"
    "github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
    "github.com/stretchr/testify/assert"
    "testing"
)

var (
    updatedTitle       = "updated stub title"
    updatedDescription = "updated stub description"
    updatedBody        = "updated stub body"
)

func Test_WhenDownstreamServicesAreSuccessful_ReturnsMappedUpdatedArticle(t *testing.T) {
    // Arrange
    request := articles.UpdateArticleServiceRequest{
        UserId:      1,
        ArticleSlug: slug.Make("testUser1 article"),
        Title:       &updatedTitle,
        Description: &updatedDescription,
        Body:        &updatedBody,
    }

    // Act
    response, err := fixture.service.UpdateArticle(fixture.ctx, &request)

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
}

func Test_WhenNoArticleIsFound_ReturnsError(t *testing.T) {
    // Arrange
    request := articles.UpdateArticleServiceRequest{
        UserId:      3,
        ArticleSlug: slug.Make("testUser2 article"),
        Title:       &updatedTitle,
        Description: &updatedDescription,
        Body:        &updatedBody,
    }

    // Act
    response, err := fixture.service.UpdateArticle(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrArticlesNotFound.Error())
}

func Test_WhenUserIdDoesNotMatch_ReturnsError(t *testing.T) {
    // Arrange
    request := articles.UpdateArticleServiceRequest{
        UserId:      2,
        ArticleSlug: slug.Make("testUser1 article"),
        Title:       &updatedTitle,
        Description: &updatedDescription,
        Body:        &updatedBody,
    }

    // Act
    response, err := fixture.service.UpdateArticle(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrArticlesNotFound.Error())
}

func Test_WhenArticleSlugExists_ReturnsError(t *testing.T) {
    // Arrange
    request := articles.UpdateArticleServiceRequest{
        UserId:      2,
        ArticleSlug: slug.Make("testUser1 article"),
        Title:       &updatedTitle,
        Description: &updatedDescription,
        Body:        &updatedBody,
    }

    // Act
    response, err := fixture.service.UpdateArticle(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrArticlesNotFound.Error())
}
