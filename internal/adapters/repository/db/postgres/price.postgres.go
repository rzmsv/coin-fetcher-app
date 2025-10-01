package postgres

import (
	"context"
	"github.com/username/coin-fetcher-app/internal/domain/interfaces"
	"github.com/username/coin-fetcher-app/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

type PriceRepository struct {
	DB *gorm.DB
}

func NewPriceRepository(db *gorm.DB) interfaces.PriceRepository {
	return &PriceRepository{
		DB: db,
	}
}

func (p *PriceRepository) Save(Coin *model.Coin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return p.DB.WithContext(ctx).Create(&Coin).Error
}

func (p *PriceRepository) GetLastPrice(coin string) (*model.Coin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var Coin *model.Coin
	err := p.DB.WithContext(ctx).Where("coin = ?", coin).Order("timestamp DESC").First(&Coin).Error
	if err != nil {
		return &model.Coin{}, err
	}
	return Coin, nil

}
func (p *PriceRepository) GetAveragePrice(since time.Time, coin string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var Result struct {
		AvgPrice float64
	}
	err := p.DB.WithContext(ctx).Model(&model.Coin{}).Select("AVG(price) as avg_price").Where("coin = ? AND timestamp >= ?", coin, since).Scan(&Result).Error
	if err != nil {
		return 0, err
	}
	return Result.AvgPrice, nil
}
