package models

import (
	"firstRest/database"
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
