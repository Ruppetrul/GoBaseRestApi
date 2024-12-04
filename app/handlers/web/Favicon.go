package web

import (
	"log"
	"net/http"
	"os"
)

func Favicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	file, err := os.ReadFile("blood_512.png")
	if err != nil {
		log.Println("Error favicon response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	_, err = w.Write(file)
	if err != nil {
		log.Println("Error robots response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
