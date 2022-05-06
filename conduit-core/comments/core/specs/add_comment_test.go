package specs

import (
    "github.com/joeymckenzie/realworld-go-kit/conduit-core/comments/domain"
    "github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_WhenDownstreamServicesAreSuccessful_ReturnsMappedCreatedComment(t *testing.T) {
    // Arrange
    request := domain.AddArticleCommentServiceRequest{
        UserId: 1,
        Body:   "stub comment body",
        Slug:   "testuser1-article",
    }

    // Act
    response, err := fixture.service.AddComment(fixture.ctx, &request)

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
    assert.IsType(t, response, &domain.CommentDto{})
}

func Test_WhenArticleDoesNotExist_ReturnsNotFound(t *testing.T) {
    // Arrange
    request := domain.AddArticleCommentServiceRequest{
        UserId: 1,
        Body:   "stub comment body",
        Slug:   "testuser4-article",
    }

    // Act
    response, err := fixture.service.AddComment(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrArticlesNotFound.Error())
}

func Test_WhenUserDoesNotExist_ReturnsNotFound(t *testing.T) {
    // Arrange
    request := domain.AddArticleCommentServiceRequest{
        UserId: 12,
        Body:   "stub comment body",
        Slug:   "testuser1-article",
    }

    // Act
    response, err := fixture.service.AddComment(fixture.ctx, &request)

    // Assert
    assert.Nil(t, response)
    assert.NotNil(t, err)
    assert.ErrorContains(t, err, utilities.ErrUserNotFound.Error())
}
