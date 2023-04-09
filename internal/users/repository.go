package users

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, username, email, password string) (*UserEntity, error)
}

func NewRepository(db *pgxpool.Pool) UsersRepository {
	return &usersRepository{
		db: db,
	}
}

type usersRepository struct {
	db *pgxpool.Pool
}

func (r *usersRepository) CreateUser(ctx context.Context, username string, email string, password string) (*UserEntity, error) {
	var user UserEntity

	if err := pgxscan.Select(ctx, r.db, &user, "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING ID"); err != nil {
		return &user, err
	}

	return &user, nil
}
