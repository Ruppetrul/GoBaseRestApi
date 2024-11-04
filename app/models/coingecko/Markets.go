package coingecko

import (
	"database/sql"
	"firstRest/database"
)

type Markets struct {
	Id           string  `db:"id"`
	Symbol       string  `db:"symbol"`
	Name         string  `db:"name"`
	CurrentPrice float32 `db:"current_price"`
}

func (p *Markets) Save() (sql.Result, error) {
	connection, err := database.GetDBInstance()
	if err != nil {
		return nil, err
	}

	return connection.Db.Exec(`INSERT INTO coingecko (id, symbol, name, current_price) 
		VALUES ($1, $2, $3, $4) ON CONFLICT (symbol) DO UPDATE
		SET current_price = EXCLUDED.current_price RETURNING symbol;`, p.Id, p.Symbol, p.Name, p.CurrentPrice)
}

func GetList() ([]Markets, error) {
	rows, err := database.Select(`SELECT id, symbol, name, current_price FROM coingecko;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markets []Markets
	for rows.Next() {
		var market Markets
		if err := rows.Scan(&market.Id, &market.Symbol, &market.Name, &market.CurrentPrice); err != nil {
			return nil, err
		}
		markets = append(markets, market)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return markets, nil
}
