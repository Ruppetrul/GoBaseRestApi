package main

import (
	"bytes"
	"encoding/json"
	"firstRest/database"
	_ "firstRest/database"
	"firstRest/front"
	"firstRest/models/coingecko"
	"firstRest/workers"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	go workers.RegisterCoinGeckoWorker()

	http.HandleFunc("/", test1)
	http.HandleFunc("/current", current)
	http.HandleFunc("/robots.txt", robots)

	env := os.Getenv("APP_ENV")

	if env == "" {
		fmt.Println("APP_ENV not found.")
	} else {
		fmt.Printf("APP_ENV: %s\n", env)
	}

	if env == "production" {
		certFile := "/etc/letsencrypt/live/crypto-visor.ru/cert.pem"
		keyFile := "/etc/letsencrypt/live/crypto-visor.ru/privkey.pem"

		if err := http.ListenAndServeTLS(":443", certFile, keyFile, nil); err != nil {
			log.Fatalf("Ошибка при запуске сервера %v", err)
		}
	} else {
		err := http.ListenAndServe(":80", nil)
		if err != nil {
			log.Fatalf("Ошибка при запуске сервера %v", err)
		}
	}
}

func current(w http.ResponseWriter, r *http.Request) {
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

func test1(w http.ResponseWriter, r *http.Request) {
	prices, err := coingecko.GetList("market_cap")

	if err != nil {
		log.Println("Error when fetching list", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	/*
		There need prepare base html and save to temp file or memory.
	*/

	index, err := template.ParseFiles("front/index.html")
	if err != nil {
		panic(err)
	}

	table, err := template.ParseFiles("front/table.html")
	if err != nil {
		panic(err)
	}

	var rows []string
	tableRow, err := template.ParseFiles("front/table_row.html")
	if err != nil {
		panic(err)
	}

	for _, v := range prices {
		var rowBuf bytes.Buffer
		if err := tableRow.Execute(&rowBuf, v); err != nil {
			panic(err)
		}
		rows = append(rows, rowBuf.String())
	}

	var tableBuf, indexBuf bytes.Buffer

	if err := table.Execute(&tableBuf, front.TableData{
		Rows: template.HTML(strings.Join(rows, "")),
	}); err != nil {
		panic(err)
	}

	tableResult := tableBuf.String()

	data := front.FrontData{
		Table: template.HTML(tableResult),
	}

	if err := index.Execute(&indexBuf, data); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(indexBuf.String())); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = database.IncrementVisitCount()
	if err != nil {
		log.Println("Error increment count response:", err)
		return
	}
}

func robots(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("User-agent: *\nAllow: /"))
	if err != nil {
		log.Println("Error robots response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
