package main

import (
	"bytes"
	"encoding/json"
	_ "firstRest/database"
	"firstRest/front"
	"firstRest/models/coingecko"
	"firstRest/workers"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", test1)
	http.HandleFunc("/current", current)
	go workers.RegisterCoinGeckoWorker()
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера %v", err)
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

	var tableBuf bytes.Buffer
	var indexBuf bytes.Buffer

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
	log.Println(indexBuf.String())
	if _, err := w.Write([]byte(indexBuf.String())); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
