package main

import (
	"encoding/json"
	_ "firstRest/database"
	"firstRest/models"
	"firstRest/workers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/current", current)
	http.HandleFunc("/test", test)
	go workers.RegisterBinanceWorker()
	//TODO "BYbit" etc. workers.
	go workers.RegisterGeneralWorker() //General calculations.
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера %v", err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	filename := "front/index.html"
	html, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprint(w, string(html))
	if err != nil {
		log.Fatal(err)
	}
}

func current(w http.ResponseWriter, r *http.Request) {
	prices, err := models.GetList()
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

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("Welcome")
}
