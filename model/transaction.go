package model

import (
	"time"
)

type Transaction struct {
    Id int32
    Amount float64
    Incoming bool
    Recurring string
    StartDate time.Time
    EndDate time.Time
}
