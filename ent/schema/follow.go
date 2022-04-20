package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Follow holds the schema definition for the Follow entity.
type Follow struct {
	ent.Schema
}

// Fields of the Follow.
func (Follow) Fields() []ent.Field {
	return []ent.Field{
		field.Int("follower_id").
			Optional(),
		field.Int("followee_id").
			Optional(),
	}
}

// Edges of the Follow.
func (Follow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_follower", User.Type).
			Ref("followers").
			Unique().
			Field("follower_id"),
		edge.From("user_followee", User.Type).
			Ref("followees").
			Unique().
			Field("followee_id"),
	}
}

func (Follow) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}
