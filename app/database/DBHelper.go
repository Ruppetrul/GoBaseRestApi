package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

type dbSingleton struct {
	Db *sql.DB
}

var instance *dbSingleton
var once sync.Once

func GetDBInstance() (*dbSingleton, error) {
	var err error

	env := os.Getenv("APP_ENV")
	var connStr string
	if env == "production" {
		connStr = "user=first_rest password=first_rest dbname=first_rest host=localhost sslmode=disable"
	} else {
		connStr = "user=first_rest password=first_rest dbname=first_rest host=postgres sslmode=disable"
	}
	once.Do(func() {
		instance = &dbSingleton{}
		instance.Db, err = sql.Open("postgres", connStr)
	})

	if err != nil {
		log.Fatalf("DB connect error: $v", err)
		return nil, err
	}

	return instance, nil
}

func Select(query string) (*sql.Rows, error) {
	connection, err := GetDBInstance()
	if err != nil {
		log.Println("Error scanning row: query", err)
	}
	return connection.Db.Query(query)
}
