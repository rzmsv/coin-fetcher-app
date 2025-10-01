package scheduler

import (
	"log"
	"time"

	"github.com/username/coin-fetcher-app/internal/application"
)

func Start(service *application.PriceService, schedulerMinute int, coin string) {

	if schedulerMinute == 0 {
		schedulerMinute = 1
	}
	ticker := time.NewTicker(time.Duration(schedulerMinute) * time.Minute)
	go func() {
		for range ticker.C {
			if err := service.FetchAndStorePriceToRedis(coin); err != nil {
				log.Printf("Scheduler error: %v", err)
			} else {
				if err := service.FetchAndStorePriceToDB(coin); err != nil {
					log.Printf("Scheduler error: %v", err)
				}
				log.Println("Price fetched and stored")
			}
		}
	}()
}
