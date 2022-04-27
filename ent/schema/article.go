package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty().
			Default(""),
		field.String("body").
			NotEmpty().
			Default(""),
		field.String("description").
			NotEmpty().
			Default(""),
		field.String("slug").
			NotEmpty().
			Unique(),
		field.Int("user_id").
			Optional(),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("articles").
			Unique().
			Field("user_id"),
		edge.To("favorites", Favorite.Type),
		edge.To("article_tags", ArticleTag.Type),
		edge.To("article_comments", Comment.Type),
	}
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
