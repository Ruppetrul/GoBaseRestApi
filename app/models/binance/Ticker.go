package binance

import (
	"firstRest/database"
	"log"
)

type Ticker struct {
	Symbol             string
	LastPrice          string
	PriceChangePercent string
	QuoteVolume        string
}

func (p *Ticker) Save() error {
	connection, err := database.GetDBInstance()

	if err != nil {
		return err
	}
	_, err = connection.Db.Exec(`INSERT INTO binance (symbol, last_price, price_change_percent, quote_volume)
		 VALUES ($1, $2, $3, $4)
		 ON CONFLICT (symbol) DO UPDATE
		 SET last_price = EXCLUDED.last_price
		 RETURNING symbol`, p.Symbol, p.LastPrice, p.PriceChangePercent, p.QuoteVolume)
	if err != nil {
		return err
	}
	return nil
}

func GetList() ([]Ticker, error) {
	prices, err := database.Select(`SELECT symbol, last_price FROM binance ORDER BY quote_volume DESC;`)

	if err != nil {
		log.Println("Error scanning row: query", err)
		return nil, err
	}

	var pricesResult []Ticker
	for prices.Next() {
		var price Ticker
		if err := prices.Scan(&price.Symbol, &price.LastPrice); err != nil {
			log.Println("Error scanning row: parse", err)
			return nil, err
		}
		pricesResult = append(pricesResult, price)
	}

	return pricesResult, nil
}
