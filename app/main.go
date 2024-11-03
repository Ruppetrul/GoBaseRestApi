package main

import (
	"encoding/json"
	_ "firstRest/database"
	"firstRest/models"
	"firstRest/models/General"
	"firstRest/workers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/current", current)
	http.HandleFunc("/test1", test1) //Read from file.
	http.HandleFunc("/test2", test2) //Read from postgres.
	http.HandleFunc("/test3", test3) //Read from memory.
	go workers.RegisterBinanceWorker()
	//TODO "BYbit" etc. workers.
	go workers.RegisterGeneralWorker() //General calculations.
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера %v", err)
	}
}

func test1(w http.ResponseWriter, r *http.Request) {
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

func test2(w http.ResponseWriter, r *http.Request) {
	html, err := General.GetFirst()
	if err != nil {
		log.Println("Error when fetching list", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(html.Html)); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func test3(w http.ResponseWriter, r *http.Request) {
	html, err := General.GetFromMemory()
	if err != nil {
		log.Println("Error when fetching list", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(html.Html)); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
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
