package core

import (
	"github.com/joeymckenzie/realworld-go-kit/ent"
)

func updateIfRequired(existingField string, requestField *string) string {
	if requestField != nil && *requestField != "" {
		return *requestField
	}

	return existingField
}

func isValidDatabaseErr(err error) bool {
	return err != nil && !ent.IsNotFound(err)
}
