package api

import (
	"encoding/json"
	"firstRest/models/coingecko"
	"log"
	"net/http"
)

func Current(w http.ResponseWriter, r *http.Request) {
	prices, err := coingecko.GetList("market_cap")
	if err != nil {
		log.Println("Error when fetching list", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(prices); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
