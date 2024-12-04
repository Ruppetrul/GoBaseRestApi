package web

import (
	"log"
	"net/http"
)

func Robots(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("User-agent: *\nAllow: /"))
	if err != nil {
		log.Println("Error robots response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
