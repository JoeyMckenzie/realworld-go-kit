package users

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Username string
	Email    string
	Password string
	Image    string
	Bio      string
}

type UsersRepository interface {
	CreateUser(ctx context.Context, username, email, password string) (*User, error)
}

func NewRepository(db *pgx.Conn) UsersRepository {
	return &usersRepository{
		db: db,
	}
}

type usersRepository struct {
	db *pgx.Conn
}

// CreateUser implements UsersRepository
func (r *usersRepository) CreateUser(ctx context.Context, username string, email string, password string) (*User, error) {
	const sql string = `
CREATE `

	panic("unimplemented")
}
