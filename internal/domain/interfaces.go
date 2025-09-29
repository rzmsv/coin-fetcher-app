package domain

import (
	"time"
)

type PriceRepository interface {
	Save(price Price) error
	GetLastPrice(symbol string) (Price, error)
	GetAveragePrice(since time.Time) (float64, error)
}

type PriceFetcher interface {
	FetchPrice(coin string) (float64, error)
}
