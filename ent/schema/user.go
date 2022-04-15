package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("username").
            NotEmpty().
            Unique(),
        field.String("email").
            NotEmpty().
            Unique(),
        field.String("password").
            NotEmpty().
            Default(""),
        field.String("bio").
            NotEmpty().
            Default(""),
        field.String("image").
            NotEmpty().
            Default(""),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("articles", Article.Type),
    }
}

func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
