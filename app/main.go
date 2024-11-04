package main

import (
	"encoding/json"
	_ "firstRest/database"
	"firstRest/models/coingecko"
	"firstRest/workers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/current", current)
	go workers.RegisterCoinGeckoWorker()
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера %v", err)
	}
}

func current(w http.ResponseWriter, r *http.Request) {
	prices, err := coingecko.GetList()
	if err != nil {
		log.Println("Error when fetching list", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println(prices)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(prices); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("Welcome")
}
