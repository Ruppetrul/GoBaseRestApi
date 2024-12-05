package coingecko

import (
	"bytes"
	"firstRest/front"
	"html/template"
	"strings"
)

func GetCurrentData() ([]byte, error) {
	prices, err := GetList("market_cap")

	if err != nil {
		return nil, err
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

	return []byte(indexBuf.String()), nil
}
