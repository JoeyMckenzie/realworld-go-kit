package articles

import (
    "context"
    "github.com/jmoiron/sqlx"
    "github.com/joeymckenzie/realworld-go-kit/internal/infrastructure/repositories"
    "golang.org/x/exp/slog"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

type articlesServiceTestFixture struct {
    ctx                 context.Context
    service             ArticlesService
    mockUsersRepository *repositories.MockUsersRepository
}

func newArticlesServiceTestFixture() *articlesServiceTestFixture {
    ctx := context.Background()
    nopLogger := slog.Default()
    dsn := os.Getenv("DSN")
    db := sqlx.MustOpen("mysql", dsn)
    articlesRepository := repositories.NewArticlesRepository(db)
    mockUsersRepository := new(repositories.MockUsersRepository)
    service := NewArticlesService(nopLogger, articlesRepository, mockUsersRepository)

    return &articlesServiceTestFixture{
        ctx:                 ctx,
        service:             service,
        mockUsersRepository: mockUsersRepository,
    }
}
