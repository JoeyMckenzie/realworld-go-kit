package shared

type UsernameContextKey struct {
    Username string
}

func GetUpdatedValueIfApplicable(requestValue string, existingValue string) string {
    if requestValue != "" {
        return requestValue
    }
    return existingValue
}
