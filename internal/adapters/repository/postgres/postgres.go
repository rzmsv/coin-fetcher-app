package postgres

import (
	"log"

	"github.com/username/coin-fetcher-app/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

var DB *gorm.DB

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	DB = db
	log.Println("Connect to Database.")
	// Auto migrate schema
	migrateDB()
	return db
}

func migrateDB() {
	DB.AutoMigrate(&domain.Price{})
	log.Println("Migration done.")
}
