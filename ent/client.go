// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/joeymckenzie/realworld-go-kit/ent/migrate"

	"github.com/joeymckenzie/realworld-go-kit/ent/article"
	"github.com/joeymckenzie/realworld-go-kit/ent/articletag"
	"github.com/joeymckenzie/realworld-go-kit/ent/favorite"
	"github.com/joeymckenzie/realworld-go-kit/ent/follow"
	"github.com/joeymckenzie/realworld-go-kit/ent/tag"
	"github.com/joeymckenzie/realworld-go-kit/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Article is the client for interacting with the Article builders.
	Article *ArticleClient
	// ArticleTag is the client for interacting with the ArticleTag builders.
	ArticleTag *ArticleTagClient
	// Favorite is the client for interacting with the Favorite builders.
	Favorite *FavoriteClient
	// Follow is the client for interacting with the Follow builders.
	Follow *FollowClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Article = NewArticleClient(c.config)
	c.ArticleTag = NewArticleTagClient(c.config)
	c.Favorite = NewFavoriteClient(c.config)
	c.Follow = NewFollowClient(c.config)
	c.Tag = NewTagClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Article:    NewArticleClient(cfg),
		ArticleTag: NewArticleTagClient(cfg),
		Favorite:   NewFavoriteClient(cfg),
		Follow:     NewFollowClient(cfg),
		Tag:        NewTagClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Article:    NewArticleClient(cfg),
		ArticleTag: NewArticleTagClient(cfg),
		Favorite:   NewFavoriteClient(cfg),
		Follow:     NewFollowClient(cfg),
		Tag:        NewTagClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Article.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Article.Use(hooks...)
	c.ArticleTag.Use(hooks...)
	c.Favorite.Use(hooks...)
	c.Follow.Use(hooks...)
	c.Tag.Use(hooks...)
	c.User.Use(hooks...)
}

// ArticleClient is a client for the Article schema.
type ArticleClient struct {
	config
}

// NewArticleClient returns a client for the Article from the given config.
func NewArticleClient(c config) *ArticleClient {
	return &ArticleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `article.Hooks(f(g(h())))`.
func (c *ArticleClient) Use(hooks ...Hook) {
	c.hooks.Article = append(c.hooks.Article, hooks...)
}

// Create returns a create builder for Article.
func (c *ArticleClient) Create() *ArticleCreate {
	mutation := newArticleMutation(c.config, OpCreate)
	return &ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Article entities.
func (c *ArticleClient) CreateBulk(builders ...*ArticleCreate) *ArticleCreateBulk {
	return &ArticleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Article.
func (c *ArticleClient) Update() *ArticleUpdate {
	mutation := newArticleMutation(c.config, OpUpdate)
	return &ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArticleClient) UpdateOne(a *Article) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticle(a))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArticleClient) UpdateOneID(id int) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticleID(id))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Article.
func (c *ArticleClient) Delete() *ArticleDelete {
	mutation := newArticleMutation(c.config, OpDelete)
	return &ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ArticleClient) DeleteOne(a *Article) *ArticleDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ArticleClient) DeleteOneID(id int) *ArticleDeleteOne {
	builder := c.Delete().Where(article.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArticleDeleteOne{builder}
}

// Query returns a query builder for Article.
func (c *ArticleClient) Query() *ArticleQuery {
	return &ArticleQuery{
		config: c.config,
	}
}

// Get returns a Article entity by its id.
func (c *ArticleClient) Get(ctx context.Context, id int) (*Article, error) {
	return c.Query().Where(article.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArticleClient) GetX(ctx context.Context, id int) *Article {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a Article.
func (c *ArticleClient) QueryAuthor(a *Article) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, article.AuthorTable, article.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFavorites queries the favorites edge of a Article.
func (c *ArticleClient) QueryFavorites(a *Article) *FavoriteQuery {
	query := &FavoriteQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(favorite.Table, favorite.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, article.FavoritesTable, article.FavoritesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryArticleTags queries the article_tags edge of a Article.
func (c *ArticleClient) QueryArticleTags(a *Article) *ArticleTagQuery {
	query := &ArticleTagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(articletag.Table, articletag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, article.ArticleTagsTable, article.ArticleTagsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArticleClient) Hooks() []Hook {
	return c.hooks.Article
}

// ArticleTagClient is a client for the ArticleTag schema.
type ArticleTagClient struct {
	config
}

// NewArticleTagClient returns a client for the ArticleTag from the given config.
func NewArticleTagClient(c config) *ArticleTagClient {
	return &ArticleTagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `articletag.Hooks(f(g(h())))`.
func (c *ArticleTagClient) Use(hooks ...Hook) {
	c.hooks.ArticleTag = append(c.hooks.ArticleTag, hooks...)
}

// Create returns a create builder for ArticleTag.
func (c *ArticleTagClient) Create() *ArticleTagCreate {
	mutation := newArticleTagMutation(c.config, OpCreate)
	return &ArticleTagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ArticleTag entities.
func (c *ArticleTagClient) CreateBulk(builders ...*ArticleTagCreate) *ArticleTagCreateBulk {
	return &ArticleTagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ArticleTag.
func (c *ArticleTagClient) Update() *ArticleTagUpdate {
	mutation := newArticleTagMutation(c.config, OpUpdate)
	return &ArticleTagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArticleTagClient) UpdateOne(at *ArticleTag) *ArticleTagUpdateOne {
	mutation := newArticleTagMutation(c.config, OpUpdateOne, withArticleTag(at))
	return &ArticleTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArticleTagClient) UpdateOneID(id int) *ArticleTagUpdateOne {
	mutation := newArticleTagMutation(c.config, OpUpdateOne, withArticleTagID(id))
	return &ArticleTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ArticleTag.
func (c *ArticleTagClient) Delete() *ArticleTagDelete {
	mutation := newArticleTagMutation(c.config, OpDelete)
	return &ArticleTagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ArticleTagClient) DeleteOne(at *ArticleTag) *ArticleTagDeleteOne {
	return c.DeleteOneID(at.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ArticleTagClient) DeleteOneID(id int) *ArticleTagDeleteOne {
	builder := c.Delete().Where(articletag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArticleTagDeleteOne{builder}
}

// Query returns a query builder for ArticleTag.
func (c *ArticleTagClient) Query() *ArticleTagQuery {
	return &ArticleTagQuery{
		config: c.config,
	}
}

// Get returns a ArticleTag entity by its id.
func (c *ArticleTagClient) Get(ctx context.Context, id int) (*ArticleTag, error) {
	return c.Query().Where(articletag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArticleTagClient) GetX(ctx context.Context, id int) *ArticleTag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryArticle queries the article edge of a ArticleTag.
func (c *ArticleTagClient) QueryArticle(at *ArticleTag) *ArticleQuery {
	query := &ArticleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := at.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(articletag.Table, articletag.FieldID, id),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, articletag.ArticleTable, articletag.ArticleColumn),
		)
		fromV = sqlgraph.Neighbors(at.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTag queries the tag edge of a ArticleTag.
func (c *ArticleTagClient) QueryTag(at *ArticleTag) *TagQuery {
	query := &TagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := at.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(articletag.Table, articletag.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, articletag.TagTable, articletag.TagColumn),
		)
		fromV = sqlgraph.Neighbors(at.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArticleTagClient) Hooks() []Hook {
	return c.hooks.ArticleTag
}

// FavoriteClient is a client for the Favorite schema.
type FavoriteClient struct {
	config
}

// NewFavoriteClient returns a client for the Favorite from the given config.
func NewFavoriteClient(c config) *FavoriteClient {
	return &FavoriteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `favorite.Hooks(f(g(h())))`.
func (c *FavoriteClient) Use(hooks ...Hook) {
	c.hooks.Favorite = append(c.hooks.Favorite, hooks...)
}

// Create returns a create builder for Favorite.
func (c *FavoriteClient) Create() *FavoriteCreate {
	mutation := newFavoriteMutation(c.config, OpCreate)
	return &FavoriteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Favorite entities.
func (c *FavoriteClient) CreateBulk(builders ...*FavoriteCreate) *FavoriteCreateBulk {
	return &FavoriteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Favorite.
func (c *FavoriteClient) Update() *FavoriteUpdate {
	mutation := newFavoriteMutation(c.config, OpUpdate)
	return &FavoriteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FavoriteClient) UpdateOne(f *Favorite) *FavoriteUpdateOne {
	mutation := newFavoriteMutation(c.config, OpUpdateOne, withFavorite(f))
	return &FavoriteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FavoriteClient) UpdateOneID(id int) *FavoriteUpdateOne {
	mutation := newFavoriteMutation(c.config, OpUpdateOne, withFavoriteID(id))
	return &FavoriteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Favorite.
func (c *FavoriteClient) Delete() *FavoriteDelete {
	mutation := newFavoriteMutation(c.config, OpDelete)
	return &FavoriteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FavoriteClient) DeleteOne(f *Favorite) *FavoriteDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FavoriteClient) DeleteOneID(id int) *FavoriteDeleteOne {
	builder := c.Delete().Where(favorite.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FavoriteDeleteOne{builder}
}

// Query returns a query builder for Favorite.
func (c *FavoriteClient) Query() *FavoriteQuery {
	return &FavoriteQuery{
		config: c.config,
	}
}

// Get returns a Favorite entity by its id.
func (c *FavoriteClient) Get(ctx context.Context, id int) (*Favorite, error) {
	return c.Query().Where(favorite.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FavoriteClient) GetX(ctx context.Context, id int) *Favorite {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryArticleFavorites queries the article_favorites edge of a Favorite.
func (c *FavoriteClient) QueryArticleFavorites(f *Favorite) *ArticleQuery {
	query := &ArticleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(favorite.Table, favorite.FieldID, id),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, favorite.ArticleFavoritesTable, favorite.ArticleFavoritesColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUserFavorites queries the user_favorites edge of a Favorite.
func (c *FavoriteClient) QueryUserFavorites(f *Favorite) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(favorite.Table, favorite.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, favorite.UserFavoritesTable, favorite.UserFavoritesColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FavoriteClient) Hooks() []Hook {
	return c.hooks.Favorite
}

// FollowClient is a client for the Follow schema.
type FollowClient struct {
	config
}

// NewFollowClient returns a client for the Follow from the given config.
func NewFollowClient(c config) *FollowClient {
	return &FollowClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `follow.Hooks(f(g(h())))`.
func (c *FollowClient) Use(hooks ...Hook) {
	c.hooks.Follow = append(c.hooks.Follow, hooks...)
}

// Create returns a create builder for Follow.
func (c *FollowClient) Create() *FollowCreate {
	mutation := newFollowMutation(c.config, OpCreate)
	return &FollowCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Follow entities.
func (c *FollowClient) CreateBulk(builders ...*FollowCreate) *FollowCreateBulk {
	return &FollowCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Follow.
func (c *FollowClient) Update() *FollowUpdate {
	mutation := newFollowMutation(c.config, OpUpdate)
	return &FollowUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FollowClient) UpdateOne(f *Follow) *FollowUpdateOne {
	mutation := newFollowMutation(c.config, OpUpdateOne, withFollow(f))
	return &FollowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FollowClient) UpdateOneID(id int) *FollowUpdateOne {
	mutation := newFollowMutation(c.config, OpUpdateOne, withFollowID(id))
	return &FollowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Follow.
func (c *FollowClient) Delete() *FollowDelete {
	mutation := newFollowMutation(c.config, OpDelete)
	return &FollowDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FollowClient) DeleteOne(f *Follow) *FollowDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FollowClient) DeleteOneID(id int) *FollowDeleteOne {
	builder := c.Delete().Where(follow.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FollowDeleteOne{builder}
}

// Query returns a query builder for Follow.
func (c *FollowClient) Query() *FollowQuery {
	return &FollowQuery{
		config: c.config,
	}
}

// Get returns a Follow entity by its id.
func (c *FollowClient) Get(ctx context.Context, id int) (*Follow, error) {
	return c.Query().Where(follow.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FollowClient) GetX(ctx context.Context, id int) *Follow {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUserFollower queries the user_follower edge of a Follow.
func (c *FollowClient) QueryUserFollower(f *Follow) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(follow.Table, follow.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, follow.UserFollowerTable, follow.UserFollowerColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUserFollowee queries the user_followee edge of a Follow.
func (c *FollowClient) QueryUserFollowee(f *Follow) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(follow.Table, follow.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, follow.UserFolloweeTable, follow.UserFolloweeColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FollowClient) Hooks() []Hook {
	return c.hooks.Follow
}

// TagClient is a client for the Tag schema.
type TagClient struct {
	config
}

// NewTagClient returns a client for the Tag from the given config.
func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tag.Hooks(f(g(h())))`.
func (c *TagClient) Use(hooks ...Hook) {
	c.hooks.Tag = append(c.hooks.Tag, hooks...)
}

// Create returns a create builder for Tag.
func (c *TagClient) Create() *TagCreate {
	mutation := newTagMutation(c.config, OpCreate)
	return &TagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tag entities.
func (c *TagClient) CreateBulk(builders ...*TagCreate) *TagCreateBulk {
	return &TagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tag.
func (c *TagClient) Update() *TagUpdate {
	mutation := newTagMutation(c.config, OpUpdate)
	return &TagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagClient) UpdateOne(t *Tag) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTag(t))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TagClient) UpdateOneID(id int) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTagID(id))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tag.
func (c *TagClient) Delete() *TagDelete {
	mutation := newTagMutation(c.config, OpDelete)
	return &TagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TagClient) DeleteOne(t *Tag) *TagDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TagClient) DeleteOneID(id int) *TagDeleteOne {
	builder := c.Delete().Where(tag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagDeleteOne{builder}
}

// Query returns a query builder for Tag.
func (c *TagClient) Query() *TagQuery {
	return &TagQuery{
		config: c.config,
	}
}

// Get returns a Tag entity by its id.
func (c *TagClient) Get(ctx context.Context, id int) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id int) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryArticleTags queries the article_tags edge of a Tag.
func (c *TagClient) QueryArticleTags(t *Tag) *ArticleTagQuery {
	query := &ArticleTagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(articletag.Table, articletag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, tag.ArticleTagsTable, tag.ArticleTagsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TagClient) Hooks() []Hook {
	return c.hooks.Tag
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryArticles queries the articles edge of a User.
func (c *UserClient) QueryArticles(u *User) *ArticleQuery {
	query := &ArticleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ArticlesTable, user.ArticlesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFavorites queries the favorites edge of a User.
func (c *UserClient) QueryFavorites(u *User) *FavoriteQuery {
	query := &FavoriteQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(favorite.Table, favorite.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.FavoritesTable, user.FavoritesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFollowers queries the followers edge of a User.
func (c *UserClient) QueryFollowers(u *User) *FollowQuery {
	query := &FollowQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(follow.Table, follow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.FollowersTable, user.FollowersColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFollowees queries the followees edge of a User.
func (c *UserClient) QueryFollowees(u *User) *FollowQuery {
	query := &FollowQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(follow.Table, follow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.FolloweesTable, user.FolloweesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
