// Code generated by entc, DO NOT EDIT.

package article

import (
	"time"
)

const (
	// Label holds the string label denoting the article type in the database.
	Label = "article"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeFavorites holds the string denoting the favorites edge name in mutations.
	EdgeFavorites = "favorites"
	// EdgeArticleTags holds the string denoting the article_tags edge name in mutations.
	EdgeArticleTags = "article_tags"
	// Table holds the table name of the article in the database.
	Table = "articles"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "articles"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_id"
	// FavoritesTable is the table that holds the favorites relation/edge.
	FavoritesTable = "favorites"
	// FavoritesInverseTable is the table name for the Favorite entity.
	// It exists in this package in order to avoid circular dependency with the "favorite" package.
	FavoritesInverseTable = "favorites"
	// FavoritesColumn is the table column denoting the favorites relation/edge.
	FavoritesColumn = "article_id"
	// ArticleTagsTable is the table that holds the article_tags relation/edge.
	ArticleTagsTable = "article_tags"
	// ArticleTagsInverseTable is the table name for the ArticleTag entity.
	// It exists in this package in order to avoid circular dependency with the "articletag" package.
	ArticleTagsInverseTable = "article_tags"
	// ArticleTagsColumn is the table column denoting the article_tags relation/edge.
	ArticleTagsColumn = "article_id"
)

// Columns holds all SQL columns for article fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldTitle,
	FieldBody,
	FieldDescription,
	FieldSlug,
	FieldUserID,
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
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultBody holds the default value on creation for the "body" field.
	DefaultBody string
	// BodyValidator is a validator for the "body" field. It is called by the builders before save.
	BodyValidator func(string) error
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
)
