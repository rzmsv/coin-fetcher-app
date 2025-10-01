package repository

import (
	"github.com/username/coin-fetcher-app/config"
	"github.com/username/coin-fetcher-app/internal/adapters/repository/db/postgres"
	"github.com/username/coin-fetcher-app/internal/adapters/repository/db/redis"
	"github.com/username/coin-fetcher-app/internal/domain/interfaces"
	"gorm.io/gorm"
)

func NewRepository(appConfig *config.AppConfig, db *gorm.DB) interfaces.Repository {
	return interfaces.Repository{
		PriceRepository: postgres.NewPriceRepository(db),
		Redis:           redis.NewRedis(appConfig),
	}
}
