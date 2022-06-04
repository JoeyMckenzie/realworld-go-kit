// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// EdgeArticles holds the string denoting the articles edge name in mutations.
	EdgeArticles = "articles"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeFavorites holds the string denoting the favorites edge name in mutations.
	EdgeFavorites = "favorites"
	// EdgeFollowers holds the string denoting the followers edge name in mutations.
	EdgeFollowers = "followers"
	// EdgeFollowees holds the string denoting the followees edge name in mutations.
	EdgeFollowees = "followees"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ArticlesTable is the table that holds the articles relation/edge.
	ArticlesTable = "articles"
	// ArticlesInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	ArticlesInverseTable = "articles"
	// ArticlesColumn is the table column denoting the articles relation/edge.
	ArticlesColumn = "user_id"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "user_id"
	// FavoritesTable is the table that holds the favorites relation/edge.
	FavoritesTable = "favorites"
	// FavoritesInverseTable is the table name for the Favorite entity.
	// It exists in this package in order to avoid circular dependency with the "favorite" package.
	FavoritesInverseTable = "favorites"
	// FavoritesColumn is the table column denoting the favorites relation/edge.
	FavoritesColumn = "user_id"
	// FollowersTable is the table that holds the followers relation/edge.
	FollowersTable = "follows"
	// FollowersInverseTable is the table name for the Follow entity.
	// It exists in this package in order to avoid circular dependency with the "follow" package.
	FollowersInverseTable = "follows"
	// FollowersColumn is the table column denoting the followers relation/edge.
	FollowersColumn = "follower_id"
	// FolloweesTable is the table that holds the followees relation/edge.
	FolloweesTable = "follows"
	// FolloweesInverseTable is the table name for the Follow entity.
	// It exists in this package in order to avoid circular dependency with the "follow" package.
	FolloweesInverseTable = "follows"
	// FolloweesColumn is the table column denoting the followees relation/edge.
	FolloweesColumn = "followee_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldUsername,
	FieldEmail,
	FieldPassword,
	FieldBio,
	FieldImage,
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
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultBio holds the default value on creation for the "bio" field.
	DefaultBio string
	// DefaultImage holds the default value on creation for the "image" field.
	DefaultImage string
)
