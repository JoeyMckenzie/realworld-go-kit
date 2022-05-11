package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/conduit-domain/comments"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WhenDownstreamServicesAreSuccessful_ReturnsNoContent(t *testing.T) {
	// Arrange
	request := comments.DeleteArticleCommentServiceRequest{
		UserId:    2,
		CommentId: 1,
	}

	// Act
	err := fixture.service.DeleteComment(fixture.ctx, &request)

	// Assert
	assert.Nil(t, err)
}

func Test_WhenCommentIsNotFound_ReturnsNotFound(t *testing.T) {
	// Arrange
	request := comments.DeleteArticleCommentServiceRequest{
		CommentId: 11,
		UserId:    2,
	}

	// Act
	err := fixture.service.DeleteComment(fixture.ctx, &request)

	// Assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, utilities.ErrCommentNotFound.Error())
}

func Test_WhenUserDoesNotOwnComment_ReturnsNotFound(t *testing.T) {
	// Arrange
	request := comments.DeleteArticleCommentServiceRequest{
		CommentId: 3,
		UserId:    2,
	}

	// Act
	err := fixture.service.DeleteComment(fixture.ctx, &request)

	// Assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, utilities.ErrCommentNotFound.Error())
}
