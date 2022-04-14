package utilities

import "database/sql"

func IsValidDbError(err error) bool {
	return err != nil && err != sql.ErrNoRows
}

func IsNotFound(err error) bool {
	return err == sql.ErrNoRows
}
