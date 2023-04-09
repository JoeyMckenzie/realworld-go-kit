package shared

import "errors"

var (
	ErrInvalidRequestBody = errors.New("unable to parse the request body")
)
