package repositories

import "firstRest/models"

type RepositoryInterface interface {
	GetCurrentPrices() []models.Price
}
