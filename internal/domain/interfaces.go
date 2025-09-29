package domain

import "time"

type PriceRepository interface {
	Save(price Price) error
	GetLastPrice() (Price, error)
	GetAveragePrice(since time.Time) (float64, error)
}

type PriceFetcher interface {
	FetchPrice() (float64, error)
}
