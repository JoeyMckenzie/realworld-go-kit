package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/conduit-core/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/conduit-shared/utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WhenDownstreamServicesAreSuccessful_ReturnsMappedCreatedArticle(t *testing.T) {
	// Arrange
	request := domain.CreateArticleServiceRequest{
		UserId:      1,
		Title:       "stub title",
		Description: "stub description",
		Body:        "stub body",
		TagList:     &[]string{"stub tag"},
	}

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &request)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
}

func Test_WhenUserDoesNotExist_ReturnsBadRequest(t *testing.T) {
	// Arrange
	request := domain.CreateArticleServiceRequest{
		UserId:      11,
		Title:       "stub article",
		Description: "stub description",
		Body:        "stub body",
		TagList:     &[]string{"stub tag"},
	}

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &request)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, utilities.ErrUserNotFound.Error())
}

func Test_WhenArticleTitleExists_ReturnsBadRequest(t *testing.T) {
	// Arrange
	request := domain.CreateArticleServiceRequest{
		UserId:      1,
		Title:       "testUser1 article",
		Description: "stub description",
		Body:        "stub body",
		TagList:     &[]string{"stub tag"},
	}

	// Act
	response, err := fixture.service.CreateArticle(fixture.ctx, &request)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, utilities.ErrArticleTitleExists.Error())
}
