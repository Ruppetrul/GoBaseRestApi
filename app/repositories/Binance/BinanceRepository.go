package Binance

import (
	"encoding/json"
	"firstRest/models"
	"fmt"
	"net/http"
)

type Repository struct {
}

func (b Repository) GetCurrentPrices() ([]models.Price, error) {
	url := Repository.GetTickerBaseUrl(b)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", response.StatusCode)
	}

	var prices []PriceResponse
	if err := json.NewDecoder(response.Body).Decode(&prices); err != nil {
		return nil, err
	}

	var priceMap []models.Price
	for _, priceResponse := range prices {
		priceMap = append(priceMap, models.Price{
			Name: priceResponse.Symbol, Price: priceResponse.Price,
		})
	}

	return priceMap, nil
}

func (Repository) GetURL() string {
	return "https://api.binance.com"
}

func (b Repository) GetTickerBaseUrl() string {
	return Repository.GetURL(b) + "/api/v3/ticker/price"
}
