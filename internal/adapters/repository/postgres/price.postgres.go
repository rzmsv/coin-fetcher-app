package postgres

import (
	"fmt"
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

func (p *PriceRepository) Save(Coin domain.Coin) error {
	return p.DB.Create(&Coin).Error
}

func (p *PriceRepository) GetLastPrice(coin string) (domain.Coin, error) {
	fmt.Println(coin)
	var Coin domain.Coin

	err := p.DB.Where("coin = ?", coin).Order("timestamp DESC").First(&Coin).Error
	if err != nil {
		return domain.Coin{}, err
	}
	return Coin, nil

}
func (p *PriceRepository) GetAveragePrice(since time.Time) (float64, error) {
	var Result struct {
		AvgPrice float64
	}
	err := p.DB.Model(&domain.Coin{}).Select("AVG(price) as avg_price").Where("timestamp = ?", since).Scan(&Result).Error
	if err != nil {
		return 0, err
	}
	return Result.AvgPrice, nil
}
