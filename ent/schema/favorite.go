package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Favorite holds the schema definition for the Favorite entity.
type Favorite struct {
	ent.Schema
}

// Fields of the Favorite.
func (Favorite) Fields() []ent.Field {
	return []ent.Field{
		field.Int("article_id").
			Optional(),
		field.Int("user_id").
			Optional(),
	}
}

// Edges of the Favorite.
func (Favorite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article_favorites", Article.Type).
			Ref("favorites").
			Unique().
			Field("article_id"),
		edge.From("user_favorites", User.Type).
			Ref("favorites").
			Unique().
			Field("user_id"),
	}
}

func (Favorite) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}
