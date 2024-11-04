package repositories

import (
	"encoding/json"
	"firstRest/models/coingecko"
	"fmt"
	"net/http"
)

type CoinGeckoRepository struct {
}

func (b CoinGeckoRepository) GetTicker() ([]coingecko.Markets, error) {
	url := markets()
	fmt.Println(url)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", response.StatusCode)
	}

	var prices []coingecko.Markets
	if err := json.NewDecoder(response.Body).Decode(&prices); err != nil {
		return nil, err
	}

	return prices, nil
}

func getURL() string  { return "https://api.coingecko.com/api" }
func markets() string { return getURL() + "/v3/coins/markets?vs_currency=usd" }
