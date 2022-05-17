// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Default: ""},
		{Name: "body", Type: field.TypeString, Default: ""},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "articles_users_articles",
				Columns:    []*schema.Column{ArticlesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ArticleTagsColumns holds the columns for the "article_tags" table.
	ArticleTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "article_id", Type: field.TypeInt, Nullable: true},
		{Name: "tag_id", Type: field.TypeInt, Nullable: true},
	}
	// ArticleTagsTable holds the schema information for the "article_tags" table.
	ArticleTagsTable = &schema.Table{
		Name:       "article_tags",
		Columns:    ArticleTagsColumns,
		PrimaryKey: []*schema.Column{ArticleTagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "article_tags_articles_article_tags",
				Columns:    []*schema.Column{ArticleTagsColumns[2]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "article_tags_tags_article_tags",
				Columns:    []*schema.Column{ArticleTagsColumns[3]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "body", Type: field.TypeString, Default: ""},
		{Name: "article_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_articles_article_comments",
				Columns:    []*schema.Column{CommentsColumns[4]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FavoritesColumns holds the columns for the "favorites" table.
	FavoritesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "article_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// FavoritesTable holds the schema information for the "favorites" table.
	FavoritesTable = &schema.Table{
		Name:       "favorites",
		Columns:    FavoritesColumns,
		PrimaryKey: []*schema.Column{FavoritesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "favorites_articles_favorites",
				Columns:    []*schema.Column{FavoritesColumns[2]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "favorites_users_favorites",
				Columns:    []*schema.Column{FavoritesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FollowsColumns holds the columns for the "follows" table.
	FollowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "follower_id", Type: field.TypeInt, Nullable: true},
		{Name: "followee_id", Type: field.TypeInt, Nullable: true},
	}
	// FollowsTable holds the schema information for the "follows" table.
	FollowsTable = &schema.Table{
		Name:       "follows",
		Columns:    FollowsColumns,
		PrimaryKey: []*schema.Column{FollowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "follows_users_followers",
				Columns:    []*schema.Column{FollowsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "follows_users_followees",
				Columns:    []*schema.Column{FollowsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "tag", Type: field.TypeString, Unique: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString, Default: ""},
		{Name: "bio", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "image", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		ArticleTagsTable,
		CommentsTable,
		FavoritesTable,
		FollowsTable,
		TagsTable,
		UsersTable,
	}
)

func init() {
	ArticlesTable.ForeignKeys[0].RefTable = UsersTable
	ArticleTagsTable.ForeignKeys[0].RefTable = ArticlesTable
	ArticleTagsTable.ForeignKeys[1].RefTable = TagsTable
	CommentsTable.ForeignKeys[0].RefTable = ArticlesTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	FavoritesTable.ForeignKeys[0].RefTable = ArticlesTable
	FavoritesTable.ForeignKeys[1].RefTable = UsersTable
	FollowsTable.ForeignKeys[0].RefTable = UsersTable
	FollowsTable.ForeignKeys[1].RefTable = UsersTable
}
