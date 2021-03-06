// Code generated by entc, DO NOT EDIT.

package tag

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// Tag applies equality check predicate on the "tag" field. It's identical to TagEQ.
func Tag(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTag), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// TagEQ applies the EQ predicate on the "tag" field.
func TagEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTag), v))
	})
}

// TagNEQ applies the NEQ predicate on the "tag" field.
func TagNEQ(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTag), v))
	})
}

// TagIn applies the In predicate on the "tag" field.
func TagIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTag), v...))
	})
}

// TagNotIn applies the NotIn predicate on the "tag" field.
func TagNotIn(vs ...string) predicate.Tag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTag), v...))
	})
}

// TagGT applies the GT predicate on the "tag" field.
func TagGT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTag), v))
	})
}

// TagGTE applies the GTE predicate on the "tag" field.
func TagGTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTag), v))
	})
}

// TagLT applies the LT predicate on the "tag" field.
func TagLT(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTag), v))
	})
}

// TagLTE applies the LTE predicate on the "tag" field.
func TagLTE(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTag), v))
	})
}

// TagContains applies the Contains predicate on the "tag" field.
func TagContains(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTag), v))
	})
}

// TagHasPrefix applies the HasPrefix predicate on the "tag" field.
func TagHasPrefix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTag), v))
	})
}

// TagHasSuffix applies the HasSuffix predicate on the "tag" field.
func TagHasSuffix(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTag), v))
	})
}

// TagEqualFold applies the EqualFold predicate on the "tag" field.
func TagEqualFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTag), v))
	})
}

// TagContainsFold applies the ContainsFold predicate on the "tag" field.
func TagContainsFold(v string) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTag), v))
	})
}

// HasArticleTags applies the HasEdge predicate on the "article_tags" edge.
func HasArticleTags() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticleTagsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ArticleTagsTable, ArticleTagsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArticleTagsWith applies the HasEdge predicate on the "article_tags" edge with a given conditions (other predicates).
func HasArticleTagsWith(preds ...predicate.ArticleTag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticleTagsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ArticleTagsTable, ArticleTagsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		p(s.Not())
	})
}
