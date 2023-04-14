package data

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type MockArticlesRepository struct {
	mock.Mock
}

func (mw *MockArticlesRepository) ResetMocks() {
	mw.Mock = mock.Mock{
		ExpectedCalls: nil,
		Calls:         nil,
	}
}

func (mw *MockArticlesRepository) BeginTransaction(ctx context.Context) (*sqlx.Tx, error) {
	//TODO implement me
	panic("implement me")
}

func (mw *MockArticlesRepository) CommitTransaction(tx *sqlx.Tx) error {
	//TODO implement me
	panic("implement me")
}

func (mw *MockArticlesRepository) CreateArticleWithTags(ctx context.Context, tx *sqlx.Tx, slug, title, description, body, string, tagList []string) (*ArticleEntity, error) {
	//TODO implement me
	panic("implement me")
}
