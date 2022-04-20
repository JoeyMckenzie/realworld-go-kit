// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/joeymckenzie/realworld-go-kit/ent/article"
	"github.com/joeymckenzie/realworld-go-kit/ent/articletag"
	"github.com/joeymckenzie/realworld-go-kit/ent/tag"
)

// ArticleTag is the model entity for the ArticleTag schema.
type ArticleTag struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// TagID holds the value of the "tag_id" field.
	TagID int `json:"tag_id,omitempty"`
	// ArticleID holds the value of the "article_id" field.
	ArticleID int `json:"article_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArticleTagQuery when eager-loading is set.
	Edges ArticleTagEdges `json:"edges"`
}

// ArticleTagEdges holds the relations/edges for other nodes in the graph.
type ArticleTagEdges struct {
	// Article holds the value of the article edge.
	Article *Article `json:"article,omitempty"`
	// Tag holds the value of the tag edge.
	Tag *Tag `json:"tag,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ArticleOrErr returns the Article value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArticleTagEdges) ArticleOrErr() (*Article, error) {
	if e.loadedTypes[0] {
		if e.Article == nil {
			// The edge article was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: article.Label}
		}
		return e.Article, nil
	}
	return nil, &NotLoadedError{edge: "article"}
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArticleTagEdges) TagOrErr() (*Tag, error) {
	if e.loadedTypes[1] {
		if e.Tag == nil {
			// The edge tag was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: tag.Label}
		}
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ArticleTag) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case articletag.FieldID, articletag.FieldTagID, articletag.FieldArticleID:
			values[i] = new(sql.NullInt64)
		case articletag.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ArticleTag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ArticleTag fields.
func (at *ArticleTag) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case articletag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			at.ID = int(value.Int64)
		case articletag.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				at.CreateTime = value.Time
			}
		case articletag.FieldTagID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tag_id", values[i])
			} else if value.Valid {
				at.TagID = int(value.Int64)
			}
		case articletag.FieldArticleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field article_id", values[i])
			} else if value.Valid {
				at.ArticleID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryArticle queries the "article" edge of the ArticleTag entity.
func (at *ArticleTag) QueryArticle() *ArticleQuery {
	return (&ArticleTagClient{config: at.config}).QueryArticle(at)
}

// QueryTag queries the "tag" edge of the ArticleTag entity.
func (at *ArticleTag) QueryTag() *TagQuery {
	return (&ArticleTagClient{config: at.config}).QueryTag(at)
}

// Update returns a builder for updating this ArticleTag.
// Note that you need to call ArticleTag.Unwrap() before calling this method if this ArticleTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *ArticleTag) Update() *ArticleTagUpdateOne {
	return (&ArticleTagClient{config: at.config}).UpdateOne(at)
}

// Unwrap unwraps the ArticleTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *ArticleTag) Unwrap() *ArticleTag {
	tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("ent: ArticleTag is not a transactional entity")
	}
	at.config.driver = tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *ArticleTag) String() string {
	var builder strings.Builder
	builder.WriteString("ArticleTag(")
	builder.WriteString(fmt.Sprintf("id=%v", at.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(at.CreateTime.Format(time.ANSIC))
	builder.WriteString(", tag_id=")
	builder.WriteString(fmt.Sprintf("%v", at.TagID))
	builder.WriteString(", article_id=")
	builder.WriteString(fmt.Sprintf("%v", at.ArticleID))
	builder.WriteByte(')')
	return builder.String()
}

// ArticleTags is a parsable slice of ArticleTag.
type ArticleTags []*ArticleTag

func (at ArticleTags) config(cfg config) {
	for _i := range at {
		at[_i].config = cfg
	}
}