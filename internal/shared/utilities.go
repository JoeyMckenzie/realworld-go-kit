package shared

import "database/sql"

func GetUpdatedValueIfApplicable(requestValue string, existingValue string) string {
    if requestValue != "" {
        return requestValue
    }
    return existingValue
}

func IsValidSqlErr(err error) bool {
    return err != nil && err != sql.ErrNoRows
}
