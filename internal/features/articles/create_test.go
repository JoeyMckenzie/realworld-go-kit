package articles

import (
    "github.com/go-faker/faker/v4"
    "github.com/google/uuid"
    "github.com/joeymckenzie/realworld-go-kit/internal/domain"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

// TODO: Swap out mocks for real PlanetScale connections
func Test_CreateReturnsSuccess_WhenDownstreamServicesAreOk(t *testing.T) {
    // Arrange
    fixture := newArticlesServiceTestFixture()

    request := domain.CreateArticleRequest{
        Article: &domain.ArticleRequest{
            Title:       faker.Sentence(),
            Description: faker.Sentence(),
            Body:        faker.Sentence(),
            TagList: []string{
                faker.Word(),
                faker.Word(),
            },
        },
    }

    fixture.mockUsersRepository.
        On("GetUserById", fixture.ctx, mock.AnythingOfType("uuid.UUID")).
        Return(&repositories.UserEntity{}, nil)

    // Act
    response, err := fixture.service.CreateArticle(fixture.ctx, request, uuid.New())

    // Assert
    assert.NotNil(t, response)
    assert.Nil(t, err)
}
