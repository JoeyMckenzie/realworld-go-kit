// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/follow"
	"github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen/ent/user"
)

// Follow is the model entity for the Follow schema.
type Follow struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// FollowerID holds the value of the "follower_id" field.
	FollowerID int `json:"follower_id,omitempty"`
	// FolloweeID holds the value of the "followee_id" field.
	FolloweeID int `json:"followee_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FollowQuery when eager-loading is set.
	Edges FollowEdges `json:"edges"`
}

// FollowEdges holds the relations/edges for other nodes in the graph.
type FollowEdges struct {
	// UserFollower holds the value of the user_follower edge.
	UserFollower *User `json:"user_follower,omitempty"`
	// UserFollowee holds the value of the user_followee edge.
	UserFollowee *User `json:"user_followee,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserFollowerOrErr returns the UserFollower value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FollowEdges) UserFollowerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.UserFollower == nil {
			// The edge user_follower was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserFollower, nil
	}
	return nil, &NotLoadedError{edge: "user_follower"}
}

// UserFolloweeOrErr returns the UserFollowee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FollowEdges) UserFolloweeOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.UserFollowee == nil {
			// The edge user_followee was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserFollowee, nil
	}
	return nil, &NotLoadedError{edge: "user_followee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Follow) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case follow.FieldID, follow.FieldFollowerID, follow.FieldFolloweeID:
			values[i] = new(sql.NullInt64)
		case follow.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Follow", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Follow fields.
func (f *Follow) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case follow.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case follow.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				f.CreateTime = value.Time
			}
		case follow.FieldFollowerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field follower_id", values[i])
			} else if value.Valid {
				f.FollowerID = int(value.Int64)
			}
		case follow.FieldFolloweeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field followee_id", values[i])
			} else if value.Valid {
				f.FolloweeID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUserFollower queries the "user_follower" edge of the Follow entity.
func (f *Follow) QueryUserFollower() *UserQuery {
	return (&FollowClient{config: f.config}).QueryUserFollower(f)
}

// QueryUserFollowee queries the "user_followee" edge of the Follow entity.
func (f *Follow) QueryUserFollowee() *UserQuery {
	return (&FollowClient{config: f.config}).QueryUserFollowee(f)
}

// Update returns a builder for updating this Follow.
// Note that you need to call Follow.Unwrap() before calling this method if this Follow
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Follow) Update() *FollowUpdateOne {
	return (&FollowClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Follow entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Follow) Unwrap() *Follow {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Follow is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Follow) String() string {
	var builder strings.Builder
	builder.WriteString("Follow(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(f.CreateTime.Format(time.ANSIC))
	builder.WriteString(", follower_id=")
	builder.WriteString(fmt.Sprintf("%v", f.FollowerID))
	builder.WriteString(", followee_id=")
	builder.WriteString(fmt.Sprintf("%v", f.FolloweeID))
	builder.WriteByte(')')
	return builder.String()
}

// Follows is a parsable slice of Follow.
type Follows []*Follow

func (f Follows) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}