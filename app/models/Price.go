package models

import (
	"firstRest/database"
	_ "github.com/lib/pq"
)

type Price struct {
	ID    int
	Name  string
	Price string
}

func (p *Price) Save() error {
	connection, err := database.GetDBInstance()

	if err != nil {
		return err
	}
	defer connection.Db.Close()

	_, err = connection.Db.Exec(`INSERT INTO prices (name, price)
		 VALUES ($1, $2)
		 ON CONFLICT (name) DO UPDATE
		 SET price = EXCLUDED.price
		 RETURNING id`, p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil
}
