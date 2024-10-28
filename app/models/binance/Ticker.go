package binance

import (
	"firstRest/database"
)

type Ticker struct {
	Symbol             string
	LastPrice          string
	PriceChangePercent string
}

func (p *Ticker) Save() error {
	connection, err := database.GetDBInstance()

	if err != nil {
		return err
	}
	_, err = connection.Db.Exec(`INSERT INTO tickers (symbol, last_price, price_change_percent)
		 VALUES ($1, $2, $3)
		 ON CONFLICT (symbol) DO UPDATE
		 SET last_price = EXCLUDED.last_price
		 RETURNING symbol`, p.Symbol, p.LastPrice, p.PriceChangePercent)
	if err != nil {
		return err
	}
	return nil
}
