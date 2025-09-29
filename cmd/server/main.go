package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/username/coin-fetcher-app/config"
	"github.com/username/coin-fetcher-app/internal/adapters/external"
	"github.com/username/coin-fetcher-app/internal/adapters/http"
	postgres "github.com/username/coin-fetcher-app/internal/adapters/repository/postgres"
	"github.com/username/coin-fetcher-app/internal/application"
	"github.com/username/coin-fetcher-app/internal/scheduler"
)

func main() {
	appConfigs := config.NewAppConfig()

	db := postgres.InitDB(appConfigs.Configs("DSN"))

	priceRepo := postgres.NewPriceRepository(db)
	priceFetcher := external.NewCoinGeckoFetcher()
	priceService := application.NewPriceService(priceRepo, priceFetcher)
	httpAdapter := http.NewEchoAdapter(priceService)
	/* -------------------------------- SCHEDULER ------------------------------- */
	schedulerTimer, err := strconv.Atoi(appConfigs.Configs("SCHEDULER_MINUTE_TIME"))
	if err != nil {
		log.Fatal("Scheduler convert time to int error! ", err)
	}
	scheduler.Start(priceService, schedulerTimer, appConfigs.Configs("COIN_FOR_SCHEDULER"))

	/* ------------------------------- START ECHO ------------------------------- */
	if err := httpAdapter.Start(fmt.Sprintf(":%s", appConfigs.Configs("APP_PORT"))); err != nil {
		log.Fatal(err)
	}
}
