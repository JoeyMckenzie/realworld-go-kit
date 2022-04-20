// Code generated by entc, DO NOT EDIT.

package favorite

import (
	"time"
)

const (
	// Label holds the string label denoting the favorite type in the database.
	Label = "favorite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldArticleID holds the string denoting the article_id field in the database.
	FieldArticleID = "article_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeArticleFavorites holds the string denoting the article_favorites edge name in mutations.
	EdgeArticleFavorites = "article_favorites"
	// EdgeUserFavorites holds the string denoting the user_favorites edge name in mutations.
	EdgeUserFavorites = "user_favorites"
	// Table holds the table name of the favorite in the database.
	Table = "favorites"
	// ArticleFavoritesTable is the table that holds the article_favorites relation/edge.
	ArticleFavoritesTable = "favorites"
	// ArticleFavoritesInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	ArticleFavoritesInverseTable = "articles"
	// ArticleFavoritesColumn is the table column denoting the article_favorites relation/edge.
	ArticleFavoritesColumn = "article_id"
	// UserFavoritesTable is the table that holds the user_favorites relation/edge.
	UserFavoritesTable = "favorites"
	// UserFavoritesInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserFavoritesInverseTable = "users"
	// UserFavoritesColumn is the table column denoting the user_favorites relation/edge.
	UserFavoritesColumn = "user_id"
)

// Columns holds all SQL columns for favorite fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldArticleID,
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
)