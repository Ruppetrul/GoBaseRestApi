package workers

import (
	"bytes"
	"firstRest/front"
	"firstRest/models"
	"firstRest/models/General"
	"firstRest/models/binance"
	"fmt"
	"text/template"
	"time"
)

func RegisterGeneralWorker() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			/*
				There need compare and analytic logic.
			*/
			list, err := binance.GetList()
			if err != nil {
				return
			}

			for _, v := range list {
				b := models.General{
					Symbol: v.Symbol, LastPrice: v.LastPrice,
					PriceChangePercent: v.PriceChangePercent, QuoteVolume: v.QuoteVolume,
				}
				err := b.Save()
				if err != nil {
					fmt.Print(err)
				}
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

			rows := ""

			var rowBuf bytes.Buffer
			tableRow, err := template.ParseFiles("front/table_row.html")
			if err != nil {
				panic(err)
			}

			for _, v := range list {
				rowBuf2 := rowBuf
				tableRow2 := tableRow
				b := models.General{
					Symbol: v.Symbol, LastPrice: v.LastPrice,
					PriceChangePercent: v.PriceChangePercent, QuoteVolume: v.QuoteVolume,
				}
				if err := tableRow2.Execute(&rowBuf2, b); err != nil {
					panic(err)
				}
				rows = rows + rowBuf2.String()
			}
			var tableBuf bytes.Buffer
			var indexBuf bytes.Buffer

			if err := table.Execute(&tableBuf, front.TableData{
				Rows: rows,
			}); err != nil {
				panic(err)
			}

			tableResult := tableBuf.String()
			// Создаем данные для замены плейсхолдеров
			data := front.FrontData{
				Table: tableResult,
			}

			if err := index.Execute(&indexBuf, data); err != nil {
				panic(err)
			}

			g := General.Html{
				Html: indexBuf.String(),
			}
			err = g.Save()
			if err != nil {
				fmt.Print(err)
			}
		}
	}
}
