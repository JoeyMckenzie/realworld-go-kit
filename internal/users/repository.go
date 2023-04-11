package users

import (
    "context"
    "github.com/google/uuid"

    "github.com/georgysavva/scany/v2/pgxscan"
    "github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository interface {
    GetUser(ctx context.Context, username, email string) (*uuid.UUID, error)
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

func (r *usersRepository) GetUser(ctx context.Context, username, email string) (*uuid.UUID, error) {
    var userId *uuid.UUID
    const query string = "SELECT id FROM users WHERE (username, email) = ($1, $2)"

    if err := r.db.QueryRow(ctx, query, username, email).Scan(userId); err != nil {
        return userId, err
    }

    return userId, nil
}

func (r *usersRepository) CreateUser(ctx context.Context, username string, email string, password string) (*UserEntity, error) {
    var user UserEntity
    const command string = "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *"

    if err := pgxscan.Get(ctx, r.db, &user, command, username, email, password); err != nil {
        return &user, err
    }

    return &user, nil
}
