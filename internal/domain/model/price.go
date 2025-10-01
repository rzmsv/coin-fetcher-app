package model

import "time"

type Coin struct {
	ID        uint
	Price     float64
	Coin      string
	Timestamp time.Time
}
