package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
	"time"
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
		log.Fatalf("DB connect error: %v", err)
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

func IncrementVisitCount(remoteAddress string) error {
	connection, err := GetDBInstance()

	currentDate := time.Now().Format("2006-01-02")

	_, err = connection.Db.Exec(`
		INSERT INTO page_visits (visit_date, visit_count, visit_ip)
		VALUES ($1, 1, $2)
		ON CONFLICT (visit_date, visit_ip)
		DO UPDATE SET visit_count = page_visits.visit_count + 1;
	`, currentDate, remoteAddress)

	return err
}
