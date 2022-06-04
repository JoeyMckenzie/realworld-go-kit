package utilities

func UpdateIfRequired(existingField string, requestField *string) string {
	if requestField != nil && *requestField != "" {
		return *requestField
	}

	return existingField
}
