package core

import (
	"github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates_GivenListWithDuplicateValues_ReturnsDedupedList(t *testing.T) {
	// Arrange
	tagList := []string{"tag", "tag", "anotherTag"}
	expected := []string{"tag", "anotherTag"}

	// Act
	result := removeDuplicates(&tagList)

	// Assert
	assert.NotEmpty(t, result)
	assert.EqualValues(t, result, expected)
}

func Test_removeDuplicates_GivenNilList_ReturnsEmptyList(t *testing.T) {
	// Arrange
	var tagList *[]string

	// Act
	result := removeDuplicates(tagList)

	// Assert
	assert.Empty(t, result)
}

func Test_removeDuplicates_GivenEmptyList_ReturnsEmptyList(t *testing.T) {
	// Arrange
	var tagList *[]string

	// Act
	result := removeDuplicates(tagList)

	// Assert
	assert.Empty(t, result)
}

func Test_containsTag_GivenNonEmptyTagWithContainingTag_ReturnsTrue(t *testing.T) {
	// Arrange
	tags := &[]persistence.TagEntity{
		*persistence.StubTag,
		*persistence.StubAnotherTag,
	}

	// Act
	result := containsTag("stub tag", tags)

	// Assert
	assert.True(t, result)
}

func Test_containsTag_GivenNonEmptyTagWithoutContainingTag_ReturnsTrue(t *testing.T) {
	// Arrange
	tags := &[]persistence.TagEntity{
		*persistence.StubTag,
		*persistence.StubAnotherTag,
	}

	// Act
	result := containsTag("some stub tag", tags)

	// Assert
	assert.False(t, result)
}

func Test_containsTag_GivenNilTags_ReturnsFalse(t *testing.T) {
	// Arrange
	var tags *[]persistence.TagEntity

	// Act
	result := containsTag("some stub tag", tags)

	// Assert
	assert.False(t, result)
}

func Test_containsTag_GivenEmptyTags_ReturnsFalse(t *testing.T) {
	// Arrange
	tags := make([]persistence.TagEntity, 0)

	// Act
	result := containsTag("some stub tag", &tags)

	// Assert
	assert.False(t, result)
}

func Test_findTag_GivenSearchValueInTags_ReturnsTag(t *testing.T) {
	// Arrange
	tags := []persistence.TagEntity{
		*persistence.StubTag,
		*persistence.StubAnotherTag,
	}

	// Assert
	result := findTag("another stub tag", &tags)

	// Act
	assert.NotNil(t, result)
	assert.EqualValues(t, tags[1], *result)
}

func Test_findTag_GivenSearchValueNotInTags_ReturnsNil(t *testing.T) {
	// Arrange
	tags := []persistence.TagEntity{
		*persistence.StubTag,
		*persistence.StubAnotherTag,
	}

	// Assert
	result := findTag("not in tag collection", &tags)

	// Act
	assert.Nil(t, result)
}
