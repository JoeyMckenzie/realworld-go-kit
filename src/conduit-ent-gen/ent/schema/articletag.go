package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// ArticleTag holds the schema definition for the ArticleTag entity.
type ArticleTag struct {
	ent.Schema
}

// Fields of the ArticleTag.
func (ArticleTag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tag_id").
			Optional(),
		field.Int("article_id").
			Optional(),
	}
}

// Edges of the ArticleTag.
func (ArticleTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).
			Ref("article_tags").
			Unique().
			Field("article_id"),
		edge.From("tag", Tag.Type).
			Ref("article_tags").
			Unique().
			Field("tag_id"),
	}
}

func (ArticleTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}
