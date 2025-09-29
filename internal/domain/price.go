package domain

import "time"

type Price struct {
	ID        uint
	Price     float64
	Symbol    string
	Timestamp time.Time
}
