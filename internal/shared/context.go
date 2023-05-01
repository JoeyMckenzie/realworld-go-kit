package shared

import (
    "github.com/gofrs/uuid"
)

type (
    TokenContextKey struct {
        UserId uuid.UUID
    }

    UsernameContextKey struct {
        Username string
    }
)
