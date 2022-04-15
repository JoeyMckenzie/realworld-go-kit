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
        field.String("slug").
            NotEmpty().
            Unique(),
    }
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("author", User.Type).
            Ref("articles").
            Unique(),
    }
}

func (Article) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}
