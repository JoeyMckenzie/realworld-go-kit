// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joeymckenzie/realworld-go-kit/ent/article"
	"github.com/joeymckenzie/realworld-go-kit/ent/articletag"
	"github.com/joeymckenzie/realworld-go-kit/ent/predicate"
	"github.com/joeymckenzie/realworld-go-kit/ent/tag"
)

// ArticleTagQuery is the builder for querying ArticleTag entities.
type ArticleTagQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ArticleTag
	// eager-loading edges.
	withArticle *ArticleQuery
	withTag     *TagQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArticleTagQuery builder.
func (atq *ArticleTagQuery) Where(ps ...predicate.ArticleTag) *ArticleTagQuery {
	atq.predicates = append(atq.predicates, ps...)
	return atq
}

// Limit adds a limit step to the query.
func (atq *ArticleTagQuery) Limit(limit int) *ArticleTagQuery {
	atq.limit = &limit
	return atq
}

// Offset adds an offset step to the query.
func (atq *ArticleTagQuery) Offset(offset int) *ArticleTagQuery {
	atq.offset = &offset
	return atq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (atq *ArticleTagQuery) Unique(unique bool) *ArticleTagQuery {
	atq.unique = &unique
	return atq
}

// Order adds an order step to the query.
func (atq *ArticleTagQuery) Order(o ...OrderFunc) *ArticleTagQuery {
	atq.order = append(atq.order, o...)
	return atq
}

// QueryArticle chains the current query on the "article" edge.
func (atq *ArticleTagQuery) QueryArticle() *ArticleQuery {
	query := &ArticleQuery{config: atq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(articletag.Table, articletag.FieldID, selector),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, articletag.ArticleTable, articletag.ArticleColumn),
		)
		fromU = sqlgraph.SetNeighbors(atq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTag chains the current query on the "tag" edge.
func (atq *ArticleTagQuery) QueryTag() *TagQuery {
	query := &TagQuery{config: atq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(articletag.Table, articletag.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, articletag.TagTable, articletag.TagColumn),
		)
		fromU = sqlgraph.SetNeighbors(atq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ArticleTag entity from the query.
// Returns a *NotFoundError when no ArticleTag was found.
func (atq *ArticleTagQuery) First(ctx context.Context) (*ArticleTag, error) {
	nodes, err := atq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{articletag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atq *ArticleTagQuery) FirstX(ctx context.Context) *ArticleTag {
	node, err := atq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ArticleTag ID from the query.
// Returns a *NotFoundError when no ArticleTag ID was found.
func (atq *ArticleTagQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = atq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{articletag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atq *ArticleTagQuery) FirstIDX(ctx context.Context) int {
	id, err := atq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ArticleTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ArticleTag entity is found.
// Returns a *NotFoundError when no ArticleTag entities are found.
func (atq *ArticleTagQuery) Only(ctx context.Context) (*ArticleTag, error) {
	nodes, err := atq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{articletag.Label}
	default:
		return nil, &NotSingularError{articletag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atq *ArticleTagQuery) OnlyX(ctx context.Context) *ArticleTag {
	node, err := atq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ArticleTag ID in the query.
// Returns a *NotSingularError when more than one ArticleTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (atq *ArticleTagQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = atq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{articletag.Label}
	default:
		err = &NotSingularError{articletag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atq *ArticleTagQuery) OnlyIDX(ctx context.Context) int {
	id, err := atq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ArticleTags.
func (atq *ArticleTagQuery) All(ctx context.Context) ([]*ArticleTag, error) {
	if err := atq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return atq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (atq *ArticleTagQuery) AllX(ctx context.Context) []*ArticleTag {
	nodes, err := atq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ArticleTag IDs.
func (atq *ArticleTagQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := atq.Select(articletag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atq *ArticleTagQuery) IDsX(ctx context.Context) []int {
	ids, err := atq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atq *ArticleTagQuery) Count(ctx context.Context) (int, error) {
	if err := atq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return atq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (atq *ArticleTagQuery) CountX(ctx context.Context) int {
	count, err := atq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atq *ArticleTagQuery) Exist(ctx context.Context) (bool, error) {
	if err := atq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return atq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (atq *ArticleTagQuery) ExistX(ctx context.Context) bool {
	exist, err := atq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArticleTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atq *ArticleTagQuery) Clone() *ArticleTagQuery {
	if atq == nil {
		return nil
	}
	return &ArticleTagQuery{
		config:      atq.config,
		limit:       atq.limit,
		offset:      atq.offset,
		order:       append([]OrderFunc{}, atq.order...),
		predicates:  append([]predicate.ArticleTag{}, atq.predicates...),
		withArticle: atq.withArticle.Clone(),
		withTag:     atq.withTag.Clone(),
		// clone intermediate query.
		sql:    atq.sql.Clone(),
		path:   atq.path,
		unique: atq.unique,
	}
}

// WithArticle tells the query-builder to eager-load the nodes that are connected to
// the "article" edge. The optional arguments are used to configure the query builder of the edge.
func (atq *ArticleTagQuery) WithArticle(opts ...func(*ArticleQuery)) *ArticleTagQuery {
	query := &ArticleQuery{config: atq.config}
	for _, opt := range opts {
		opt(query)
	}
	atq.withArticle = query
	return atq
}

// WithTag tells the query-builder to eager-load the nodes that are connected to
// the "tag" edge. The optional arguments are used to configure the query builder of the edge.
func (atq *ArticleTagQuery) WithTag(opts ...func(*TagQuery)) *ArticleTagQuery {
	query := &TagQuery{config: atq.config}
	for _, opt := range opts {
		opt(query)
	}
	atq.withTag = query
	return atq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ArticleTag.Query().
//		GroupBy(articletag.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (atq *ArticleTagQuery) GroupBy(field string, fields ...string) *ArticleTagGroupBy {
	group := &ArticleTagGroupBy{config: atq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := atq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return atq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.ArticleTag.Query().
//		Select(articletag.FieldCreateTime).
//		Scan(ctx, &v)
//
func (atq *ArticleTagQuery) Select(fields ...string) *ArticleTagSelect {
	atq.fields = append(atq.fields, fields...)
	return &ArticleTagSelect{ArticleTagQuery: atq}
}

func (atq *ArticleTagQuery) prepareQuery(ctx context.Context) error {
	for _, f := range atq.fields {
		if !articletag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if atq.path != nil {
		prev, err := atq.path(ctx)
		if err != nil {
			return err
		}
		atq.sql = prev
	}
	return nil
}

func (atq *ArticleTagQuery) sqlAll(ctx context.Context) ([]*ArticleTag, error) {
	var (
		nodes       = []*ArticleTag{}
		_spec       = atq.querySpec()
		loadedTypes = [2]bool{
			atq.withArticle != nil,
			atq.withTag != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ArticleTag{config: atq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, atq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := atq.withArticle; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ArticleTag)
		for i := range nodes {
			fk := nodes[i].ArticleID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(article.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "article_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Article = n
			}
		}
	}

	if query := atq.withTag; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ArticleTag)
		for i := range nodes {
			fk := nodes[i].TagID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(tag.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "tag_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Tag = n
			}
		}
	}

	return nodes, nil
}

func (atq *ArticleTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atq.querySpec()
	_spec.Node.Columns = atq.fields
	if len(atq.fields) > 0 {
		_spec.Unique = atq.unique != nil && *atq.unique
	}
	return sqlgraph.CountNodes(ctx, atq.driver, _spec)
}

func (atq *ArticleTagQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := atq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (atq *ArticleTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   articletag.Table,
			Columns: articletag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: articletag.FieldID,
			},
		},
		From:   atq.sql,
		Unique: true,
	}
	if unique := atq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := atq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, articletag.FieldID)
		for i := range fields {
			if fields[i] != articletag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := atq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (atq *ArticleTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(atq.driver.Dialect())
	t1 := builder.Table(articletag.Table)
	columns := atq.fields
	if len(columns) == 0 {
		columns = articletag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if atq.sql != nil {
		selector = atq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if atq.unique != nil && *atq.unique {
		selector.Distinct()
	}
	for _, p := range atq.predicates {
		p(selector)
	}
	for _, p := range atq.order {
		p(selector)
	}
	if offset := atq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArticleTagGroupBy is the group-by builder for ArticleTag entities.
type ArticleTagGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atgb *ArticleTagGroupBy) Aggregate(fns ...AggregateFunc) *ArticleTagGroupBy {
	atgb.fns = append(atgb.fns, fns...)
	return atgb
}

// Scan applies the group-by query and scans the result into the given value.
func (atgb *ArticleTagGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := atgb.path(ctx)
	if err != nil {
		return err
	}
	atgb.sql = query
	return atgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (atgb *ArticleTagGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := atgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (atgb *ArticleTagGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(atgb.fields) > 1 {
		return nil, errors.New("ent: ArticleTagGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := atgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (atgb *ArticleTagGroupBy) StringsX(ctx context.Context) []string {
	v, err := atgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atgb *ArticleTagGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = atgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{articletag.Label}
	default:
		err = fmt.Errorf("ent: ArticleTagGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (atgb *ArticleTagGroupBy) StringX(ctx context.Context) string {
	v, err := atgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (atgb *ArticleTagGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(atgb.fields) > 1 {
		return nil, errors.New("ent: ArticleTagGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := atgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (atgb *ArticleTagGroupBy) IntsX(ctx context.Context) []int {
	v, err := atgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atgb *ArticleTagGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = atgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{articletag.Label}
	default:
		err = fmt.Errorf("ent: ArticleTagGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (atgb *ArticleTagGroupBy) IntX(ctx context.Context) int {
	v, err := atgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (atgb *ArticleTagGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(atgb.fields) > 1 {
		return nil, errors.New("ent: ArticleTagGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := atgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (atgb *ArticleTagGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := atgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atgb *ArticleTagGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = atgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{articletag.Label}
	default:
		err = fmt.Errorf("ent: ArticleTagGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (atgb *ArticleTagGroupBy) Float64X(ctx context.Context) float64 {
	v, err := atgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (atgb *ArticleTagGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(atgb.fields) > 1 {
		return nil, errors.New("ent: ArticleTagGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := atgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (atgb *ArticleTagGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := atgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (atgb *ArticleTagGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = atgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{articletag.Label}
	default:
		err = fmt.Errorf("ent: ArticleTagGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (atgb *ArticleTagGroupBy) BoolX(ctx context.Context) bool {
	v, err := atgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (atgb *ArticleTagGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range atgb.fields {
		if !articletag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := atgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (atgb *ArticleTagGroupBy) sqlQuery() *sql.Selector {
	selector := atgb.sql.Select()
	aggregation := make([]string, 0, len(atgb.fns))
	for _, fn := range atgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(atgb.fields)+len(atgb.fns))
		for _, f := range atgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(atgb.fields...)...)
}

// ArticleTagSelect is the builder for selecting fields of ArticleTag entities.
type ArticleTagSelect struct {
	*ArticleTagQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ats *ArticleTagSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ats.prepareQuery(ctx); err != nil {
		return err
	}
	ats.sql = ats.ArticleTagQuery.sqlQuery(ctx)
	return ats.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ats *ArticleTagSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ats.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ats *ArticleTagSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ats.fields) > 1 {
		return nil, errors.New("ent: ArticleTagSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ats *ArticleTagSelect) StringsX(ctx context.Context) []string {
	v, err := ats.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ats *ArticleTagSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ats.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{articletag.Label}
	default:
		err = fmt.Errorf("ent: ArticleTagSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ats *ArticleTagSelect) StringX(ctx context.Context) string {
	v, err := ats.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ats *ArticleTagSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ats.fields) > 1 {
		return nil, errors.New("ent: ArticleTagSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ats *ArticleTagSelect) IntsX(ctx context.Context) []int {
	v, err := ats.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ats *ArticleTagSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ats.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{articletag.Label}
	default:
		err = fmt.Errorf("ent: ArticleTagSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ats *ArticleTagSelect) IntX(ctx context.Context) int {
	v, err := ats.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ats *ArticleTagSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ats.fields) > 1 {
		return nil, errors.New("ent: ArticleTagSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ats *ArticleTagSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ats.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ats *ArticleTagSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ats.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{articletag.Label}
	default:
		err = fmt.Errorf("ent: ArticleTagSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ats *ArticleTagSelect) Float64X(ctx context.Context) float64 {
	v, err := ats.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ats *ArticleTagSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ats.fields) > 1 {
		return nil, errors.New("ent: ArticleTagSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ats *ArticleTagSelect) BoolsX(ctx context.Context) []bool {
	v, err := ats.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ats *ArticleTagSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ats.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{articletag.Label}
	default:
		err = fmt.Errorf("ent: ArticleTagSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ats *ArticleTagSelect) BoolX(ctx context.Context) bool {
	v, err := ats.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ats *ArticleTagSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ats.sql.Query()
	if err := ats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}