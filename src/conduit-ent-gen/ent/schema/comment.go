package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("body").
			NotEmpty().
			Default(""),
		field.Int("article_id").
			Optional(),
		field.Int("user_id").
			Optional(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).
			Ref("article_comments").
			Unique().
			Field("article_id"),
		edge.From("user", User.Type).
			Ref("comments").
			Unique().
			Field("user_id"),
	}
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
