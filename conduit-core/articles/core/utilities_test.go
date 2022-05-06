package core

import (
	"github.com/joeymckenzie/realworld-go-kit/ent"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var tags = []*ent.Tag{
	{
		ID:         1,
		CreateTime: time.Now(),
		Tag:        "stub tag",
	},
	{
		ID:         1,
		CreateTime: time.Now(),
		Tag:        "another stub tag",
	},
}

func Test_removeDuplicates_GivenListWithDuplicateValues_ReturnsDedupedList(t *testing.T) {
	// Arrange
	tagList := []string{"tag", "tag", "anotherTag"}
	expected := []string{"tag", "anotherTag"}

	// Act
	result := removeDuplicates(&tagList)

	// Assert
	assert.NotEmpty(t, result)
	assert.Contains(t, result, expected[0], expected[1])
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
	// Act
	result := containsTag("stub tag", tags)

	// Assert
	assert.True(t, result)
}

func Test_containsTag_GivenNonEmptyTagWithoutContainingTag_ReturnsTrue(t *testing.T) {
	// Act
	result := containsTag("some stub tag", tags)

	// Assert
	assert.False(t, result)
}

func Test_containsTag_GivenEmptyTags_ReturnsFalse(t *testing.T) {
	// Arrange
	tags := make([]*ent.Tag, 0)

	// Act
	result := containsTag("some stub tag", tags)

	// Assert
	assert.False(t, result)
}

func Test_findTag_GivenSearchValueInTags_ReturnsTag(t *testing.T) {
	// Act
	result := firstOrDefaultTag("another stub tag", tags)

	// Assert
	assert.NotNil(t, result)
	assert.EqualValues(t, tags[1], result)
}

func Test_findTag_GivenSearchValueNotInTags_ReturnsNil(t *testing.T) {
	// Act
	result := firstOrDefaultTag("not in tag collection", tags)

	// Assert
	assert.Nil(t, result)
}
