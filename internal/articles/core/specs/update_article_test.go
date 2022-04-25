package specs

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/domain"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WhenDownstreamServicesAreSuccessful_ReturnsMappedUpdatedArticle(t *testing.T) {
	// Arrange
	request := domain.UpsertArticleServiceRequest{
		UserId:      1,
		ArticleId:   2,
		Title:       "updated stub title",
		Description: "updated stub description",
		Body:        "updated stub body",
		TagList:     &[]string{"updated stub tag"},
	}

	// Act
	response, err := fixture.service.UpdateArticle(fixture.ctx, &request)

	// Assert
	assert.NotNil(t, response)
	assert.Nil(t, err)
}

func Test_WhenNoArticleIsFound_ReturnsError(t *testing.T) {
	// Arrange
	request := domain.UpsertArticleServiceRequest{
		UserId:      2,
		ArticleId:   22,
		Title:       "updated stub title",
		Description: "updated stub description",
		Body:        "updated stub body",
		TagList:     &[]string{"updated stub tag"},
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
	request := domain.UpsertArticleServiceRequest{
		UserId:      3,
		ArticleId:   1,
		Title:       "updated stub title",
		Description: "updated stub description",
		Body:        "updated stub body",
		TagList:     &[]string{"updated stub tag"},
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
	request := domain.UpsertArticleServiceRequest{
		UserId:      2,
		ArticleId:   3,
		Title:       "testUser1 article",
		Description: "updated stub description",
		Body:        "updated stub body",
		TagList:     &[]string{"updated stub tag"},
	}

	// Act
	response, err := fixture.service.UpdateArticle(fixture.ctx, &request)

	// Assert
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, utilities.ErrArticleTitleExists.Error())
}
