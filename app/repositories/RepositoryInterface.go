package repositories

import "firstRest/models"

type Interface interface {
	GetCurrentPrices() []models.Price
}
