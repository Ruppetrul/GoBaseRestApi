package main

import (
	"firstRest/models/coingecko"
	"log"
	"net/http"
)

func current(w http.ResponseWriter, r *http.Request) {
	data, err := coingecko.GetCurrentData()

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Println("GetCurrentData error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err != nil {
		log.Println("Error increment count response:", err)
		return
	}
}
