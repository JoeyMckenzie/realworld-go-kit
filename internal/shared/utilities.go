package shared

func GetUpdatedValueIfApplicable(requestValue string, existingValue string) string {
	if requestValue != "" {
		return requestValue
	}
	return existingValue
}
