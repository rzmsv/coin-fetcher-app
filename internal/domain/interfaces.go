package domain

import (
	"time"
)

type PriceRepository interface {
	Save(price Coin) error
	GetLastPrice(symbol string) (Coin, error)
	GetAveragePrice(since time.Time, coin string) (float64, error)
}

type PriceFetcher interface {
	FetchPrice(coin string) (float64, error)
}
