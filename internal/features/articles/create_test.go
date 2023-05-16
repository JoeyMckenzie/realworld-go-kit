package articles

import (
	"github.com/go-faker/faker/v4"
	"github.com/gofrs/uuid"
	"github.com/joeymckenzie/realworld-go-kit/internal/domain"
	"github.com/joeymckenzie/realworld-go-kit/internal/shared"
	"github.com/stretchr/testify/assert"
)

func (s *ArticlesServiceTestSuite) Test_ReturnsSuccess_WhenDownstreamServicesAreOk() {
	// Arrange
	request := domain.CreateArticleRequest{
		Article: &domain.CreateArticleRequestDto{
			Title:       faker.Sentence(),
			Description: faker.Sentence(),
			Body:        faker.Sentence(),
			TagList: []string{
				faker.Word(),
				faker.Word(),
			},
		},
	}

	// Act
	response, err := s.Service.CreateArticle(s.Ctx, request, s.SeedUserId)

	// Assert
	assert.NotNil(s.T(), response)
	assert.Nil(s.T(), err)
}

func (s *ArticlesServiceTestSuite) Test_ReturnsError_WhenUserIsNotFound() {
	request := domain.CreateArticleRequest{
		Article: &domain.CreateArticleRequestDto{
			Title:       faker.Sentence(),
			Description: faker.Sentence(),
			Body:        faker.Sentence(),
			TagList: []string{
				faker.Word(),
				faker.Word(),
			},
		},
	}

	// Act
	response, err := s.Service.CreateArticle(s.Ctx, request, uuid.Must(uuid.NewV4()))

	// Assert
	assert.Equal(s.T(), &domain.Article{}, response)
	assert.Error(s.T(), err)
	assert.ErrorIs(s.T(), err.(*shared.ApiError[string]).Err, shared.ErrUserNotFound)
}
