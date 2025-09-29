package external

import (
	"context"
	"fmt"
	"net/http"
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
	fmt.Println(resp)
	return 0, err
}
