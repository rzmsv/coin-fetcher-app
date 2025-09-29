package scheduler

import (
	"github.com/username/coin-fetcher-app/internal/application"
	"log"
	"time"
)

const coin = "ethereum"

func Start(service *application.PriceService) {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			if err := service.FetchAndStorePrice(coin); err != nil {
				log.Printf("Scheduler error: %v", err)
			} else {
				log.Println("Price fetched and stored")
			}
		}
	}()
}
