// Code generated by entc, DO NOT EDIT.

package favorite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// ArticleID applies equality check predicate on the "article_id" field. It's identical to ArticleIDEQ.
func ArticleID(v int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArticleID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Favorite {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Favorite(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.Favorite {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Favorite(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// ArticleIDEQ applies the EQ predicate on the "article_id" field.
func ArticleIDEQ(v int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArticleID), v))
	})
}

// ArticleIDNEQ applies the NEQ predicate on the "article_id" field.
func ArticleIDNEQ(v int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldArticleID), v))
	})
}

// ArticleIDIn applies the In predicate on the "article_id" field.
func ArticleIDIn(vs ...int) predicate.Favorite {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Favorite(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldArticleID), v...))
	})
}

// ArticleIDNotIn applies the NotIn predicate on the "article_id" field.
func ArticleIDNotIn(vs ...int) predicate.Favorite {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Favorite(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldArticleID), v...))
	})
}

// ArticleIDIsNil applies the IsNil predicate on the "article_id" field.
func ArticleIDIsNil() predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldArticleID)))
	})
}

// ArticleIDNotNil applies the NotNil predicate on the "article_id" field.
func ArticleIDNotNil() predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldArticleID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Favorite {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Favorite(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Favorite {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Favorite(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// HasArticleFavorites applies the HasEdge predicate on the "article_favorites" edge.
func HasArticleFavorites() predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticleFavoritesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ArticleFavoritesTable, ArticleFavoritesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArticleFavoritesWith applies the HasEdge predicate on the "article_favorites" edge with a given conditions (other predicates).
func HasArticleFavoritesWith(preds ...predicate.Article) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticleFavoritesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ArticleFavoritesTable, ArticleFavoritesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserFavorites applies the HasEdge predicate on the "user_favorites" edge.
func HasUserFavorites() predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserFavoritesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserFavoritesTable, UserFavoritesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserFavoritesWith applies the HasEdge predicate on the "user_favorites" edge with a given conditions (other predicates).
func HasUserFavoritesWith(preds ...predicate.User) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserFavoritesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserFavoritesTable, UserFavoritesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Favorite) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Favorite) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
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
func Not(p predicate.Favorite) predicate.Favorite {
	return predicate.Favorite(func(s *sql.Selector) {
		p(s.Not())
	})
}
