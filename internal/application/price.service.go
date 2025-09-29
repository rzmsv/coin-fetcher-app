package application

import (
	"fmt"
	"time"

	"github.com/username/coin-fetcher-app/internal/domain"
)

type PriceService struct {
	repo    domain.PriceRepository
	fetcher domain.PriceFetcher
}

func NewPriceService(repo domain.PriceRepository, fetcher domain.PriceFetcher) *PriceService {
	return &PriceService{repo, fetcher}
}

func (s *PriceService) FetchAndStorePrice(coin string) error {
	price, err := s.fetcher.FetchPrice(coin)
	if err != nil {
		return err
	}
	return s.repo.Save(domain.Price{Price: price, Timestamp: time.Now()})
}

func (s *PriceService) GetLastPrice(symbol string) (domain.Price, error) {
	return s.repo.GetLastPrice(symbol)
}

func (s *PriceService) GetAveragePrice(interval string) (float64, error) {
	var since time.Time
	switch interval {
	case "1min":
		since = time.Now().Add(-1 * time.Minute)
	case "5min":
		since = time.Now().Add(-5 * time.Minute)
	case "1day":
		since = time.Now().Add(-24 * time.Hour)
	default:
		return 0, fmt.Errorf("invalid interval")
	}
	return s.repo.GetAveragePrice(since)
}
