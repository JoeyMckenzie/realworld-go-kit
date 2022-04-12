package core

import "github.com/joeymckenzie/realworld-go-kit/internal/articles/persistence"

func removeDuplicates(tags *[]string) []string {
	if tags == nil || len(*tags) == 0 {
		return []string{}
	}

	var dedupedTags []string
	tagCountMap := make(map[string]int)

	for _, tag := range *tags {
		if _, exists := tagCountMap[tag]; !exists {
			tagCountMap[tag] = 1
		}
	}

	for tag, _ := range tagCountMap {
		dedupedTags = append(dedupedTags, tag)
	}

	return dedupedTags
}

func containsTag(searchValue string, tags *[]persistence.TagEntity) bool {
	if tags == nil || len(*tags) == 0 {
		return false
	}

	for _, value := range *tags {
		if value.Tag == searchValue {
			return true
		}
	}

	return false
}

func findTag(searchTag string, tags *[]persistence.TagEntity) *persistence.TagEntity {
	if !containsTag(searchTag, tags) {
		return nil
	}

	for _, tag := range *tags {
		if tag.Tag == searchTag {
			return &tag
		}
	}

	return nil
}
