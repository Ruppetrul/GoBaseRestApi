package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Price struct {
	ID    int
	Name  string
	Price string
}

func (p *Price) Save() error {
	connStr := "user=first_rest password=first_rest dbname=first_rest host=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO prices (name, price)
		 VALUES ($1, $2)
		 ON CONFLICT (name) DO UPDATE
		 SET price = EXCLUDED.price
		 RETURNING id`, p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil
}
