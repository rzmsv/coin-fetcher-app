package main

import (
	"log"

	"github.com/username/coin-fetcher-app/config"
	"github.com/username/coin-fetcher-app/internal/adapters/external"
	"github.com/username/coin-fetcher-app/internal/adapters/http"
	postgres "github.com/username/coin-fetcher-app/internal/adapters/repository/postgres"
	"github.com/username/coin-fetcher-app/internal/application"
)

func main() {
	appConfigs := config.NewAppConfig()

	db := postgres.InitDB(appConfigs.Configs("DSN"))

	priceRepo := postgres.NewPriceRepository(db)
	priceFetcher := external.NewCoinGeckoFetcher()
	priceService := application.NewPriceService(priceRepo, priceFetcher)
	httpAdapter := http.NewEchoAdapter(priceService)
	if err := httpAdapter.Start(appConfigs.Configs("APP_PORT")); err != nil {
		log.Fatal(err)
	}
}
