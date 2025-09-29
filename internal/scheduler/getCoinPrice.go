package scheduler

import (
	"github.com/username/coin-fetcher-app/internal/application"
	"log"
	"time"
)

func Start(service *application.PriceService, schedulerMinute int, coin string) {
	if schedulerMinute == 0 {
		schedulerMinute = 1
	}
	ticker := time.NewTicker(time.Duration(schedulerMinute) * time.Minute)
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
