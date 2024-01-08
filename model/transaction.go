package model

import (
	"time"
)

type Transaction struct {
    Amount float64
    Incoming bool
    Recurring string
    StartDate time.Time
    EndDate time.Time
}
