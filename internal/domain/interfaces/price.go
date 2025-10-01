package interfaces

import (
	model "github.com/username/coin-fetcher-app/internal/domain/model"
	"time"
)

type PriceRepository interface {
	Save(price *model.Coin) error
	GetLastPrice(symbol string) (*model.Coin, error)
	GetAveragePrice(since time.Time, coin string) (float64, error)
}

type PriceFetcher interface {
	FetchPrice(coin string) (float64, error)
}
