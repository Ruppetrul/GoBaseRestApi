package main

import (
	"bytes"
	"firstRest/database"
	"firstRest/front"
	"firstRest/models/coingecko"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func current(w http.ResponseWriter, r *http.Request) {
	prices, err := coingecko.GetList("market_cap")

	if err != nil {
		log.Println("Error when fetching list", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	/*
		There need prepare base html and save to temp file or memory.
	*/
	index := template.Must(template.ParseFiles("front/index.html"))
	table := template.Must(template.ParseFiles("front/table.html"))
	tableRow := template.Must(template.ParseFiles("front/table_row.html"))

	var rows []string
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
