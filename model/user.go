package model

import "github.com/google/uuid"

type User struct {
    Id uuid.UUID
    Username string
    Email string
}
