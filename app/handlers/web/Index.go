package web

import (
	"firstRest/database"
	"firstRest/models/coingecko"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data, err := coingecko.GetCurrentData()
	if err != nil {
		log.Println("GetCurrentData error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = database.IncrementVisitCount(r.RemoteAddr)
	if err != nil {
		log.Println("Error increment count response:", err)
		return
	}
}
