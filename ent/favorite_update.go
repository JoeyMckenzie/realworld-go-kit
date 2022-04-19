// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joeymckenzie/realworld-go-kit/ent/article"
	"github.com/joeymckenzie/realworld-go-kit/ent/favorite"
	"github.com/joeymckenzie/realworld-go-kit/ent/predicate"
	"github.com/joeymckenzie/realworld-go-kit/ent/user"
)

// FavoriteUpdate is the builder for updating Favorite entities.
type FavoriteUpdate struct {
	config
	hooks    []Hook
	mutation *FavoriteMutation
}

// Where appends a list predicates to the FavoriteUpdate builder.
func (fu *FavoriteUpdate) Where(ps ...predicate.Favorite) *FavoriteUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetArticleID sets the "article_id" field.
func (fu *FavoriteUpdate) SetArticleID(i int) *FavoriteUpdate {
	fu.mutation.SetArticleID(i)
	return fu
}

// SetNillableArticleID sets the "article_id" field if the given value is not nil.
func (fu *FavoriteUpdate) SetNillableArticleID(i *int) *FavoriteUpdate {
	if i != nil {
		fu.SetArticleID(*i)
	}
	return fu
}

// ClearArticleID clears the value of the "article_id" field.
func (fu *FavoriteUpdate) ClearArticleID() *FavoriteUpdate {
	fu.mutation.ClearArticleID()
	return fu
}

// SetUserID sets the "user_id" field.
func (fu *FavoriteUpdate) SetUserID(i int) *FavoriteUpdate {
	fu.mutation.SetUserID(i)
	return fu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fu *FavoriteUpdate) SetNillableUserID(i *int) *FavoriteUpdate {
	if i != nil {
		fu.SetUserID(*i)
	}
	return fu
}

// ClearUserID clears the value of the "user_id" field.
func (fu *FavoriteUpdate) ClearUserID() *FavoriteUpdate {
	fu.mutation.ClearUserID()
	return fu
}

// SetArticleFavoritesID sets the "article_favorites" edge to the Article entity by ID.
func (fu *FavoriteUpdate) SetArticleFavoritesID(id int) *FavoriteUpdate {
	fu.mutation.SetArticleFavoritesID(id)
	return fu
}

// SetNillableArticleFavoritesID sets the "article_favorites" edge to the Article entity by ID if the given value is not nil.
func (fu *FavoriteUpdate) SetNillableArticleFavoritesID(id *int) *FavoriteUpdate {
	if id != nil {
		fu = fu.SetArticleFavoritesID(*id)
	}
	return fu
}

// SetArticleFavorites sets the "article_favorites" edge to the Article entity.
func (fu *FavoriteUpdate) SetArticleFavorites(a *Article) *FavoriteUpdate {
	return fu.SetArticleFavoritesID(a.ID)
}

// SetUserFavoritesID sets the "user_favorites" edge to the User entity by ID.
func (fu *FavoriteUpdate) SetUserFavoritesID(id int) *FavoriteUpdate {
	fu.mutation.SetUserFavoritesID(id)
	return fu
}

// SetNillableUserFavoritesID sets the "user_favorites" edge to the User entity by ID if the given value is not nil.
func (fu *FavoriteUpdate) SetNillableUserFavoritesID(id *int) *FavoriteUpdate {
	if id != nil {
		fu = fu.SetUserFavoritesID(*id)
	}
	return fu
}

// SetUserFavorites sets the "user_favorites" edge to the User entity.
func (fu *FavoriteUpdate) SetUserFavorites(u *User) *FavoriteUpdate {
	return fu.SetUserFavoritesID(u.ID)
}

// Mutation returns the FavoriteMutation object of the builder.
func (fu *FavoriteUpdate) Mutation() *FavoriteMutation {
	return fu.mutation
}

// ClearArticleFavorites clears the "article_favorites" edge to the Article entity.
func (fu *FavoriteUpdate) ClearArticleFavorites() *FavoriteUpdate {
	fu.mutation.ClearArticleFavorites()
	return fu
}

// ClearUserFavorites clears the "user_favorites" edge to the User entity.
func (fu *FavoriteUpdate) ClearUserFavorites() *FavoriteUpdate {
	fu.mutation.ClearUserFavorites()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FavoriteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FavoriteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FavoriteUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FavoriteUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FavoriteUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FavoriteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   favorite.Table,
			Columns: favorite.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: favorite.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fu.mutation.ArticleFavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favorite.ArticleFavoritesTable,
			Columns: []string{favorite.ArticleFavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ArticleFavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favorite.ArticleFavoritesTable,
			Columns: []string{favorite.ArticleFavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.UserFavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favorite.UserFavoritesTable,
			Columns: []string{favorite.UserFavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.UserFavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favorite.UserFavoritesTable,
			Columns: []string{favorite.UserFavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{favorite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FavoriteUpdateOne is the builder for updating a single Favorite entity.
type FavoriteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FavoriteMutation
}

// SetArticleID sets the "article_id" field.
func (fuo *FavoriteUpdateOne) SetArticleID(i int) *FavoriteUpdateOne {
	fuo.mutation.SetArticleID(i)
	return fuo
}

// SetNillableArticleID sets the "article_id" field if the given value is not nil.
func (fuo *FavoriteUpdateOne) SetNillableArticleID(i *int) *FavoriteUpdateOne {
	if i != nil {
		fuo.SetArticleID(*i)
	}
	return fuo
}

// ClearArticleID clears the value of the "article_id" field.
func (fuo *FavoriteUpdateOne) ClearArticleID() *FavoriteUpdateOne {
	fuo.mutation.ClearArticleID()
	return fuo
}

// SetUserID sets the "user_id" field.
func (fuo *FavoriteUpdateOne) SetUserID(i int) *FavoriteUpdateOne {
	fuo.mutation.SetUserID(i)
	return fuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fuo *FavoriteUpdateOne) SetNillableUserID(i *int) *FavoriteUpdateOne {
	if i != nil {
		fuo.SetUserID(*i)
	}
	return fuo
}

// ClearUserID clears the value of the "user_id" field.
func (fuo *FavoriteUpdateOne) ClearUserID() *FavoriteUpdateOne {
	fuo.mutation.ClearUserID()
	return fuo
}

// SetArticleFavoritesID sets the "article_favorites" edge to the Article entity by ID.
func (fuo *FavoriteUpdateOne) SetArticleFavoritesID(id int) *FavoriteUpdateOne {
	fuo.mutation.SetArticleFavoritesID(id)
	return fuo
}

// SetNillableArticleFavoritesID sets the "article_favorites" edge to the Article entity by ID if the given value is not nil.
func (fuo *FavoriteUpdateOne) SetNillableArticleFavoritesID(id *int) *FavoriteUpdateOne {
	if id != nil {
		fuo = fuo.SetArticleFavoritesID(*id)
	}
	return fuo
}

// SetArticleFavorites sets the "article_favorites" edge to the Article entity.
func (fuo *FavoriteUpdateOne) SetArticleFavorites(a *Article) *FavoriteUpdateOne {
	return fuo.SetArticleFavoritesID(a.ID)
}

// SetUserFavoritesID sets the "user_favorites" edge to the User entity by ID.
func (fuo *FavoriteUpdateOne) SetUserFavoritesID(id int) *FavoriteUpdateOne {
	fuo.mutation.SetUserFavoritesID(id)
	return fuo
}

// SetNillableUserFavoritesID sets the "user_favorites" edge to the User entity by ID if the given value is not nil.
func (fuo *FavoriteUpdateOne) SetNillableUserFavoritesID(id *int) *FavoriteUpdateOne {
	if id != nil {
		fuo = fuo.SetUserFavoritesID(*id)
	}
	return fuo
}

// SetUserFavorites sets the "user_favorites" edge to the User entity.
func (fuo *FavoriteUpdateOne) SetUserFavorites(u *User) *FavoriteUpdateOne {
	return fuo.SetUserFavoritesID(u.ID)
}

// Mutation returns the FavoriteMutation object of the builder.
func (fuo *FavoriteUpdateOne) Mutation() *FavoriteMutation {
	return fuo.mutation
}

// ClearArticleFavorites clears the "article_favorites" edge to the Article entity.
func (fuo *FavoriteUpdateOne) ClearArticleFavorites() *FavoriteUpdateOne {
	fuo.mutation.ClearArticleFavorites()
	return fuo
}

// ClearUserFavorites clears the "user_favorites" edge to the User entity.
func (fuo *FavoriteUpdateOne) ClearUserFavorites() *FavoriteUpdateOne {
	fuo.mutation.ClearUserFavorites()
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FavoriteUpdateOne) Select(field string, fields ...string) *FavoriteUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Favorite entity.
func (fuo *FavoriteUpdateOne) Save(ctx context.Context) (*Favorite, error) {
	var (
		err  error
		node *Favorite
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FavoriteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FavoriteUpdateOne) SaveX(ctx context.Context) *Favorite {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FavoriteUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FavoriteUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FavoriteUpdateOne) sqlSave(ctx context.Context) (_node *Favorite, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   favorite.Table,
			Columns: favorite.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: favorite.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Favorite.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, favorite.FieldID)
		for _, f := range fields {
			if !favorite.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != favorite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fuo.mutation.ArticleFavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favorite.ArticleFavoritesTable,
			Columns: []string{favorite.ArticleFavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ArticleFavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favorite.ArticleFavoritesTable,
			Columns: []string{favorite.ArticleFavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.UserFavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favorite.UserFavoritesTable,
			Columns: []string{favorite.UserFavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.UserFavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favorite.UserFavoritesTable,
			Columns: []string{favorite.UserFavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Favorite{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{favorite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
