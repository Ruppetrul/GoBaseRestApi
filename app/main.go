package main

import (
	"encoding/json"
	"firstRest/models"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/current", current)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера %v", err)
	}
}

func current(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	prices := []models.Price{
		{ID: 1, Name: "BTC", Price: float32(rand.Intn(60000) + 10000)},
		{ID: 1, Name: "LTC", Price: float32(rand.Intn(300) + 50)},
		{ID: 1, Name: "DOGE", Price: float32(rand.Intn(0.3*100) + 0.1*100)},
	}
	_ = json.NewEncoder(w).Encode(prices)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("Welcome")
}
