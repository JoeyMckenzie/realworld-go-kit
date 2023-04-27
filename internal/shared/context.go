package shared

import (
    "github.com/google/uuid"
)

type (
    TokenContextKey struct {
        UserId uuid.UUID
    }

    UsernameContextKey struct {
        Username string
    }
)
