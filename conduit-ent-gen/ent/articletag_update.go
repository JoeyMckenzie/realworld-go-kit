// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/article"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/articletag"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/predicate"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/tag"
)

// ArticleTagUpdate is the builder for updating ArticleTag entities.
type ArticleTagUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleTagMutation
}

// Where appends a list predicates to the ArticleTagUpdate builder.
func (atu *ArticleTagUpdate) Where(ps ...predicate.ArticleTag) *ArticleTagUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetTagID sets the "tag_id" field.
func (atu *ArticleTagUpdate) SetTagID(i int) *ArticleTagUpdate {
	atu.mutation.SetTagID(i)
	return atu
}

// SetNillableTagID sets the "tag_id" field if the given value is not nil.
func (atu *ArticleTagUpdate) SetNillableTagID(i *int) *ArticleTagUpdate {
	if i != nil {
		atu.SetTagID(*i)
	}
	return atu
}

// ClearTagID clears the value of the "tag_id" field.
func (atu *ArticleTagUpdate) ClearTagID() *ArticleTagUpdate {
	atu.mutation.ClearTagID()
	return atu
}

// SetArticleID sets the "article_id" field.
func (atu *ArticleTagUpdate) SetArticleID(i int) *ArticleTagUpdate {
	atu.mutation.SetArticleID(i)
	return atu
}

// SetNillableArticleID sets the "article_id" field if the given value is not nil.
func (atu *ArticleTagUpdate) SetNillableArticleID(i *int) *ArticleTagUpdate {
	if i != nil {
		atu.SetArticleID(*i)
	}
	return atu
}

// ClearArticleID clears the value of the "article_id" field.
func (atu *ArticleTagUpdate) ClearArticleID() *ArticleTagUpdate {
	atu.mutation.ClearArticleID()
	return atu
}

// SetArticle sets the "article" edge to the Article entity.
func (atu *ArticleTagUpdate) SetArticle(a *Article) *ArticleTagUpdate {
	return atu.SetArticleID(a.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (atu *ArticleTagUpdate) SetTag(t *Tag) *ArticleTagUpdate {
	return atu.SetTagID(t.ID)
}

// Mutation returns the ArticleTagMutation object of the builder.
func (atu *ArticleTagUpdate) Mutation() *ArticleTagMutation {
	return atu.mutation
}

// ClearArticle clears the "article" edge to the Article entity.
func (atu *ArticleTagUpdate) ClearArticle() *ArticleTagUpdate {
	atu.mutation.ClearArticle()
	return atu
}

// ClearTag clears the "tag" edge to the Tag entity.
func (atu *ArticleTagUpdate) ClearTag() *ArticleTagUpdate {
	atu.mutation.ClearTag()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *ArticleTagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atu.hooks) == 0 {
		affected, err = atu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atu.mutation = mutation
			affected, err = atu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atu.hooks) - 1; i >= 0; i-- {
			if atu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atu *ArticleTagUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *ArticleTagUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *ArticleTagUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (atu *ArticleTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   articletag.Table,
			Columns: articletag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: articletag.FieldID,
			},
		},
	}
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atu.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.ArticleTable,
			Columns: []string{articletag.ArticleColumn},
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
	if nodes := atu.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.ArticleTable,
			Columns: []string{articletag.ArticleColumn},
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
	if atu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.TagTable,
			Columns: []string{articletag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.TagTable,
			Columns: []string{articletag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{articletag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ArticleTagUpdateOne is the builder for updating a single ArticleTag entity.
type ArticleTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArticleTagMutation
}

// SetTagID sets the "tag_id" field.
func (atuo *ArticleTagUpdateOne) SetTagID(i int) *ArticleTagUpdateOne {
	atuo.mutation.SetTagID(i)
	return atuo
}

// SetNillableTagID sets the "tag_id" field if the given value is not nil.
func (atuo *ArticleTagUpdateOne) SetNillableTagID(i *int) *ArticleTagUpdateOne {
	if i != nil {
		atuo.SetTagID(*i)
	}
	return atuo
}

// ClearTagID clears the value of the "tag_id" field.
func (atuo *ArticleTagUpdateOne) ClearTagID() *ArticleTagUpdateOne {
	atuo.mutation.ClearTagID()
	return atuo
}

// SetArticleID sets the "article_id" field.
func (atuo *ArticleTagUpdateOne) SetArticleID(i int) *ArticleTagUpdateOne {
	atuo.mutation.SetArticleID(i)
	return atuo
}

// SetNillableArticleID sets the "article_id" field if the given value is not nil.
func (atuo *ArticleTagUpdateOne) SetNillableArticleID(i *int) *ArticleTagUpdateOne {
	if i != nil {
		atuo.SetArticleID(*i)
	}
	return atuo
}

// ClearArticleID clears the value of the "article_id" field.
func (atuo *ArticleTagUpdateOne) ClearArticleID() *ArticleTagUpdateOne {
	atuo.mutation.ClearArticleID()
	return atuo
}

// SetArticle sets the "article" edge to the Article entity.
func (atuo *ArticleTagUpdateOne) SetArticle(a *Article) *ArticleTagUpdateOne {
	return atuo.SetArticleID(a.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (atuo *ArticleTagUpdateOne) SetTag(t *Tag) *ArticleTagUpdateOne {
	return atuo.SetTagID(t.ID)
}

// Mutation returns the ArticleTagMutation object of the builder.
func (atuo *ArticleTagUpdateOne) Mutation() *ArticleTagMutation {
	return atuo.mutation
}

// ClearArticle clears the "article" edge to the Article entity.
func (atuo *ArticleTagUpdateOne) ClearArticle() *ArticleTagUpdateOne {
	atuo.mutation.ClearArticle()
	return atuo
}

// ClearTag clears the "tag" edge to the Tag entity.
func (atuo *ArticleTagUpdateOne) ClearTag() *ArticleTagUpdateOne {
	atuo.mutation.ClearTag()
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *ArticleTagUpdateOne) Select(field string, fields ...string) *ArticleTagUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated ArticleTag entity.
func (atuo *ArticleTagUpdateOne) Save(ctx context.Context) (*ArticleTag, error) {
	var (
		err  error
		node *ArticleTag
	)
	if len(atuo.hooks) == 0 {
		node, err = atuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atuo.mutation = mutation
			node, err = atuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atuo.hooks) - 1; i >= 0; i-- {
			if atuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *ArticleTagUpdateOne) SaveX(ctx context.Context) *ArticleTag {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *ArticleTagUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *ArticleTagUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (atuo *ArticleTagUpdateOne) sqlSave(ctx context.Context) (_node *ArticleTag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   articletag.Table,
			Columns: articletag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: articletag.FieldID,
			},
		},
	}
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ArticleTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, articletag.FieldID)
		for _, f := range fields {
			if !articletag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != articletag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atuo.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.ArticleTable,
			Columns: []string{articletag.ArticleColumn},
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
	if nodes := atuo.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.ArticleTable,
			Columns: []string{articletag.ArticleColumn},
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
	if atuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.TagTable,
			Columns: []string{articletag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.TagTable,
			Columns: []string{articletag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ArticleTag{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{articletag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
