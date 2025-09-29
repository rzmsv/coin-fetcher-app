package postgres

import (
	"time"

	"github.com/username/coin-fetcher-app/internal/domain"
	"gorm.io/gorm"
)

type PriceRepository struct {
	DB *gorm.DB
}

func NewPriceRepository(db *gorm.DB) domain.PriceRepository {
	return &PriceRepository{
		DB: db,
	}
}

func (p *PriceRepository) Save(price domain.Price) error {
	return p.DB.Create(&price).Error
}
func (p *PriceRepository) GetLastPrice(symbol string) (domain.Price, error) {
	var Price domain.Price
	err := p.DB.First(&Price).Error
	if err != nil {
		return domain.Price{}, err
	}
	return Price, nil

}
func (p *PriceRepository) GetAveragePrice(since time.Time) (float64, error) {
	var Result struct {
		AvgPrice float64
	}
	err := p.DB.Model(&domain.Price{}).Select("AVG(price) as avg_price").Where("timestamp = ?", since).Scan(&Result).Error
	if err != nil {
		return 0, err
	}
	return Result.AvgPrice, nil
}
