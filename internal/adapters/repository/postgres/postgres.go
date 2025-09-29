package postgres

import (
	"github.com/username/coin-fetcher-app/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres(dsn string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	/* --------------------------------- MIGRATE -------------------------------- */
	db.AutoMigrate(&domain.Price{})
	/* --------------------------------- MIGRATE -------------------------------- */

	return &Postgres{db}, nil
}
