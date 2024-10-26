package repositories

import (
	"firstRest/models"
	"math/rand"
)

type BinanceRepository struct {
}

func (b BinanceRepository) GetCurrentPrices() []models.Price {
	return []models.Price{
		{ID: 1, Name: "BTC", Price: float32(rand.Intn(60000) + 10000)},
		{ID: 1, Name: "LTC", Price: float32(rand.Intn(300) + 50)},
		{ID: 1, Name: "DOGE", Price: float32(rand.Intn(0.3*100) + 0.1*100)},
	}
}
