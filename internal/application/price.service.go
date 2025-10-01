package application

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	interfaces "github.com/username/coin-fetcher-app/internal/domain/interfaces"
	model "github.com/username/coin-fetcher-app/internal/domain/model"
)

type PriceService struct {
	repository interfaces.Repository
	fetcher    interfaces.PriceFetcher
}

func NewPriceService(repository interfaces.Repository, fetcher interfaces.PriceFetcher) *PriceService {
	return &PriceService{repository, fetcher}
}

func (s *PriceService) FetchAndStorePriceToRedis(coin string) error {
	price, err := s.fetcher.FetchPrice(coin)
	if err != nil {
		return err
	}
	return s.repository.Redis.Set(coin, &model.Coin{Coin: strings.ToLower(coin), Price: price, Timestamp: time.Now()}, 60*time.Second)
	// return s.repository.Save(&model.Coin{Coin: strings.ToLower(coin), Price: price, Timestamp: time.Now()})
}

func (s *PriceService) FetchAndStorePriceToDB(coin string) error {
	var data *model.Coin
	dataFromRedis, err := s.repository.Redis.Get(coin)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(dataFromRedis), &data)
	return s.repository.Save(data)
}

func (s *PriceService) GetLastPrice(coin string) (*model.Coin, error) {
	return s.repository.GetLastPrice(coin)
}

func (s *PriceService) GetAveragePrice(interval string, coin string) (float64, error) {
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
	return s.repository.GetAveragePrice(since, coin)
}
