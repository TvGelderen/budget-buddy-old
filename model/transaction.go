package model

import "time"

type Transaction struct {
    Amount float64
    Incoming bool
    Recurring string
    Date time.Time
    NextDate time.Time
    EndDate time.Time
}
