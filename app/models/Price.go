package models

import (
	"firstRest/database"
	"log"
)

type Price struct {
	Name  string
	Price string
}

func (p *Price) Save() error {
	connection, err := database.GetDBInstance()

	if err != nil {
		return err
	}

	_, err = connection.Db.Exec(`INSERT INTO prices (name, price)
		 VALUES ($1, $2)
		 ON CONFLICT (name) DO UPDATE
		 SET price = EXCLUDED.price
		 RETURNING name`, p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil
}

func GetList() ([]Price, error) {
	prices, err := database.Select(`SELECT name, price FROM prices;`)

	if err != nil {
		log.Println("Error scanning row: query", err)
		return nil, err
	}

	var pricesResult []Price
	for prices.Next() {
		var price Price
		if err := prices.Scan(&price.Name, &price.Price); err != nil {
			log.Println("Error scanning row: parse", err)
			return nil, err
		}
		pricesResult = append(pricesResult, price)
	}

	return pricesResult, nil
}
