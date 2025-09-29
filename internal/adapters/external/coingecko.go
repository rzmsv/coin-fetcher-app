package external

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type CoinGeckoFetcher struct{}

func NewCoinGeckoFetcher() *CoinGeckoFetcher {
	return &CoinGeckoFetcher{}
}

func (f *CoinGeckoFetcher) FetchPrice(coin string) (float64, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := http.Get(fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", coin))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var data map[string]struct {
		USD float64 `json:"usd"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	coin = strings.ToLower(coin)
	priceData, exists := data[coin]
	if !exists {
		return 0, fmt.Errorf("price data for coin %s not found", coin)
	}

	return priceData.USD, err
}
