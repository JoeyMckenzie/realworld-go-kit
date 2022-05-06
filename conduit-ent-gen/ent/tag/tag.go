// Code generated by entc, DO NOT EDIT.

package tag

import (
	"time"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// EdgeArticleTags holds the string denoting the article_tags edge name in mutations.
	EdgeArticleTags = "article_tags"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// ArticleTagsTable is the table that holds the article_tags relation/edge.
	ArticleTagsTable = "article_tags"
	// ArticleTagsInverseTable is the table name for the ArticleTag entity.
	// It exists in this package in order to avoid circular dependency with the "articletag" package.
	ArticleTagsInverseTable = "article_tags"
	// ArticleTagsColumn is the table column denoting the article_tags relation/edge.
	ArticleTagsColumn = "tag_id"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldTag,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// TagValidator is a validator for the "tag" field. It is called by the builders before save.
	TagValidator func(string) error
)