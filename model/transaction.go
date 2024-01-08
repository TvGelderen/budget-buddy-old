package model

import (
	"database/sql"
	"time"
)

type Transaction struct {
    Amount float64
    Incoming bool
    Recurring string
    Date time.Time
    NextDate sql.NullTime
    EndDate sql.NullTime
}
