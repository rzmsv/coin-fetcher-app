package postgres

import (
	"time"

	"github.com/username/coin-fetcher-app/internal/domain"
)

func (p *Postgres) Save(price *domain.Price) error {
	return p.DB.Create(&price).Error
}
func (p *Postgres) GetLastPrice(symbol string) (*domain.Price, error) {
	var Price *domain.Price
	err := p.DB.First(&Price).Error
	if err != nil {
		return &domain.Price{}, err
	}
	return Price, nil

}
func (p *Postgres) GetAveragePrice(since time.Time) (float64, error) {
	var Result struct {
		AvgPrice float64
	}
	err := p.DB.Model(&domain.Price{}).Select("AVG(price) as avg_price").Where("timestamp = ?", since).Scan(&Result).Error
	if err != nil {
		return 0, err
	}
	return Result.AvgPrice, nil
}
