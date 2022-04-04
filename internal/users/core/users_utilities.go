package core

import (
	"database/sql"
	"github.com/joeymckenzie/realworld-go-kit/pkg/utilities"
)

func updateIfRequired(existingField string, requestField *string) string {
	if requestField != nil && *requestField != "" {
		return *requestField
	}

	return existingField
}

func isValidDatabaseErr(err error) bool {
	return err != nil && err != sql.ErrNoRows
}

func handleRepositoryErrors(err error) error {
	if err != nil && err != sql.ErrNoRows {
		return err
	} else if err == sql.ErrNoRows {
		return utilities.ErrUserNotFound
	}

	return nil
}
