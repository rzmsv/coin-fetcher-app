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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Channel to receive response or error
	type result struct {
		price float64
		err   error
	}
	resultChan := make(chan result, 1)

	go func() {
		resp, err := http.Get(fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", coin))
		if err != nil {
			resultChan <- result{0, err}
			return
		}
		defer resp.Body.Close()
		var data map[string]struct {
			USD float64 `json:"usd"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			resultChan <- result{0, err}
			return
		}
		coin = strings.ToLower(coin)
		priceData, exists := data[coin]
		if !exists {
			resultChan <- result{0, fmt.Errorf("price data for coin %s not found", coin)}
			return
		}
		resultChan <- result{priceData.USD, nil}
	}()

	select {
	case res := <-resultChan:
		if res.err != nil {
			return 0, res.err
		}
		fmt.Printf("Fetched price for %s: %f USD\n", coin, res.price)
		return res.price, nil
	case <-ctx.Done():
		return 0, fmt.Errorf("request timed out or was canceled: %v", ctx.Err())
	}
}
