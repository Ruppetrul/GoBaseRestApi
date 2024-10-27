package main

import (
	"database/sql"
	"encoding/json"
	"firstRest/workers"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/current", current)
	go workers.RegisterCurrentPriceWorker()
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера %v", err)
	}
}

func current(w http.ResponseWriter, r *http.Request) {
	connStr := "user=first_rest password=first_rest dbname=first_rest host=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	prices, err := db.Query(`SELECT name, price FROM prices;`)
	if err != nil {
		log.Println("Error scanning row: query", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var pricesResult []PriceResponse
	for prices.Next() {
		var price PriceResponse
		if err := prices.Scan(&price.Name, &price.Price); err != nil {
			log.Println("Error scanning row: parse", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		pricesResult = append(pricesResult, price)
	}

	if err := prices.Err(); err != nil {
		log.Println("Error scanning row: parse 2", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pricesResult); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("Welcome")
}
