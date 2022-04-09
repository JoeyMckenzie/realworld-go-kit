package core

import (
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
