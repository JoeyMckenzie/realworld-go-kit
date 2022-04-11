package core

import (
	"database/sql"
)

func updateIfRequired(existingField string, requestField *string) string {
	if requestField != nil && *requestField != "" {
		return *requestField
	}

	return existingField
}

func isValidDatabaseErr(err error) bool {
	return err != nil && !isNotFound(err)
}

func isNotFound(err error) bool {
	return err == sql.ErrNoRows
}
