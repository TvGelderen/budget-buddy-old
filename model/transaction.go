package model

import (
	"time"
)

type Transaction struct {
    Id int32
    Amount float64
    Incoming bool
    Description string
    Recurring string
    StartDate time.Time
    EndDate time.Time
    Date time.Time
}
