package coingecko

import (
	"database/sql"
	"firstRest/database"
	"fmt"
	"strconv"
	"strings"
)

type Markets struct {
	Id                    string  `db:"id"`
	Symbol                string  `db:"symbol"`
	Name                  string  `db:"name"`
	CurrentPrice          float32 `json:"current_price" db:"current_price"`
	FormattedPrice        string
	MarketCap             int64 `json:"market_cap" db:"market_cap"`
	FormattedCap          string
	PriceChangePercentage float32 `json:"price_change_percentage_24h" db:"price_change_percentage_24h"`
}

func FormatPrice(price float32) string {
	priceStr := strconv.FormatFloat(float64(price), 'f', -1, 32)

	parts := strings.Split(priceStr, ".")

	intPart := parts[0]
	intPartWithCommas := ""

	for i, digit := range intPart {
		if i > 0 && (len(intPart)-i)%3 == 0 {
			intPartWithCommas += "."
		}
		intPartWithCommas += string(digit)
	}

	if len(parts) == 1 {
		return intPartWithCommas
	}

	return intPartWithCommas + "," + parts[1]
}

func (p *Markets) Save() (sql.Result, error) {
	connection, err := database.GetDBInstance()
	if err != nil {
		return nil, err
	}

	return connection.Db.Exec(`INSERT INTO coingecko (id, symbol, name, current_price, market_cap, price_change_percentage_24h) 
		VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (symbol) DO UPDATE
		SET current_price = EXCLUDED.current_price RETURNING symbol;`,
		p.Id, p.Symbol, p.Name, p.CurrentPrice, p.MarketCap, p.PriceChangePercentage)
}

func GetList(orderField string) ([]Markets, error) {
	rows, err := database.Select(fmt.Sprintf("SELECT "+
		"id, UPPER(symbol) as symbol, name,"+
		" ROUND(current_price, 8) AS current_price, market_cap, ROUND(price_change_percentage_24h, 2)"+
		" FROM coingecko ORDER BY %s DESC;", orderField))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markets []Markets
	for rows.Next() {
		var market Markets
		if err := rows.Scan(
			&market.Id, &market.Symbol, &market.Name,
			&market.CurrentPrice, &market.MarketCap, &market.PriceChangePercentage,
		); err != nil {
			return nil, err
		}
		market.FormattedPrice = FormatPrice(market.CurrentPrice)
		market.FormattedCap = FormatPrice(float32(market.MarketCap))
		markets = append(markets, market)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return markets, nil
}
