// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/article"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/articletag"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/comment"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/favorite"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/user"
)

// ArticleCreate is the builder for creating a Article entity.
type ArticleCreate struct {
	config
	mutation *ArticleMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ac *ArticleCreate) SetCreateTime(t time.Time) *ArticleCreate {
	ac.mutation.SetCreateTime(t)
	return ac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableCreateTime(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetCreateTime(*t)
	}
	return ac
}

// SetUpdateTime sets the "update_time" field.
func (ac *ArticleCreate) SetUpdateTime(t time.Time) *ArticleCreate {
	ac.mutation.SetUpdateTime(t)
	return ac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableUpdateTime(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetUpdateTime(*t)
	}
	return ac
}

// SetTitle sets the "title" field.
func (ac *ArticleCreate) SetTitle(s string) *ArticleCreate {
	ac.mutation.SetTitle(s)
	return ac
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableTitle(s *string) *ArticleCreate {
	if s != nil {
		ac.SetTitle(*s)
	}
	return ac
}

// SetBody sets the "body" field.
func (ac *ArticleCreate) SetBody(s string) *ArticleCreate {
	ac.mutation.SetBody(s)
	return ac
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableBody(s *string) *ArticleCreate {
	if s != nil {
		ac.SetBody(*s)
	}
	return ac
}

// SetDescription sets the "description" field.
func (ac *ArticleCreate) SetDescription(s string) *ArticleCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableDescription(s *string) *ArticleCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetSlug sets the "slug" field.
func (ac *ArticleCreate) SetSlug(s string) *ArticleCreate {
	ac.mutation.SetSlug(s)
	return ac
}

// SetUserID sets the "user_id" field.
func (ac *ArticleCreate) SetUserID(i int) *ArticleCreate {
	ac.mutation.SetUserID(i)
	return ac
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableUserID(i *int) *ArticleCreate {
	if i != nil {
		ac.SetUserID(*i)
	}
	return ac
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (ac *ArticleCreate) SetAuthorID(id int) *ArticleCreate {
	ac.mutation.SetAuthorID(id)
	return ac
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (ac *ArticleCreate) SetNillableAuthorID(id *int) *ArticleCreate {
	if id != nil {
		ac = ac.SetAuthorID(*id)
	}
	return ac
}

// SetAuthor sets the "author" edge to the User entity.
func (ac *ArticleCreate) SetAuthor(u *User) *ArticleCreate {
	return ac.SetAuthorID(u.ID)
}

// AddFavoriteIDs adds the "favorites" edge to the Favorite entity by IDs.
func (ac *ArticleCreate) AddFavoriteIDs(ids ...int) *ArticleCreate {
	ac.mutation.AddFavoriteIDs(ids...)
	return ac
}

// AddFavorites adds the "favorites" edges to the Favorite entity.
func (ac *ArticleCreate) AddFavorites(f ...*Favorite) *ArticleCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ac.AddFavoriteIDs(ids...)
}

// AddArticleTagIDs adds the "article_tags" edge to the ArticleTag entity by IDs.
func (ac *ArticleCreate) AddArticleTagIDs(ids ...int) *ArticleCreate {
	ac.mutation.AddArticleTagIDs(ids...)
	return ac
}

// AddArticleTags adds the "article_tags" edges to the ArticleTag entity.
func (ac *ArticleCreate) AddArticleTags(a ...*ArticleTag) *ArticleCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddArticleTagIDs(ids...)
}

// AddArticleCommentIDs adds the "article_comments" edge to the Comment entity by IDs.
func (ac *ArticleCreate) AddArticleCommentIDs(ids ...int) *ArticleCreate {
	ac.mutation.AddArticleCommentIDs(ids...)
	return ac
}

// AddArticleComments adds the "article_comments" edges to the Comment entity.
func (ac *ArticleCreate) AddArticleComments(c ...*Comment) *ArticleCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ac.AddArticleCommentIDs(ids...)
}

// Mutation returns the ArticleMutation object of the builder.
func (ac *ArticleCreate) Mutation() *ArticleMutation {
	return ac.mutation
}

// Save creates the Article in the database.
func (ac *ArticleCreate) Save(ctx context.Context) (*Article, error) {
	var (
		err  error
		node *Article
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArticleCreate) SaveX(ctx context.Context) *Article {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArticleCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArticleCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArticleCreate) defaults() {
	if _, ok := ac.mutation.CreateTime(); !ok {
		v := article.DefaultCreateTime()
		ac.mutation.SetCreateTime(v)
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		v := article.DefaultUpdateTime()
		ac.mutation.SetUpdateTime(v)
	}
	if _, ok := ac.mutation.Title(); !ok {
		v := article.DefaultTitle
		ac.mutation.SetTitle(v)
	}
	if _, ok := ac.mutation.Body(); !ok {
		v := article.DefaultBody
		ac.mutation.SetBody(v)
	}
	if _, ok := ac.mutation.Description(); !ok {
		v := article.DefaultDescription
		ac.mutation.SetDescription(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArticleCreate) check() error {
	if _, ok := ac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Article.create_time"`)}
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Article.update_time"`)}
	}
	if _, ok := ac.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Article.title"`)}
	}
	if v, ok := ac.mutation.Title(); ok {
		if err := article.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Article.title": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Article.body"`)}
	}
	if v, ok := ac.mutation.Body(); ok {
		if err := article.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Article.body": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Article.description"`)}
	}
	if v, ok := ac.mutation.Description(); ok {
		if err := article.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Article.description": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Article.slug"`)}
	}
	if v, ok := ac.mutation.Slug(); ok {
		if err := article.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Article.slug": %w`, err)}
		}
	}
	return nil
}

func (ac *ArticleCreate) sqlSave(ctx context.Context) (*Article, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *ArticleCreate) createSpec() (*Article, *sqlgraph.CreateSpec) {
	var (
		_node = &Article{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: article.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: article.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: article.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ac.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: article.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ac.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := ac.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldBody,
		})
		_node.Body = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := ac.mutation.Slug(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldSlug,
		})
		_node.Slug = value
	}
	if nodes := ac.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   article.AuthorTable,
			Columns: []string{article.AuthorColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.FavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.FavoritesTable,
			Columns: []string{article.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: favorite.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ArticleTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.ArticleTagsTable,
			Columns: []string{article.ArticleTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: articletag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ArticleCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.ArticleCommentsTable,
			Columns: []string{article.ArticleCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArticleCreateBulk is the builder for creating many Article entities in bulk.
type ArticleCreateBulk struct {
	config
	builders []*ArticleCreate
}

// Save creates the Article entities in the database.
func (acb *ArticleCreateBulk) Save(ctx context.Context) ([]*Article, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Article, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArticleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArticleCreateBulk) SaveX(ctx context.Context) []*Article {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArticleCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArticleCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
