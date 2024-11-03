package workers

import (
	"bytes"
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

			var buf bytes.Buffer

			// Выполняем шаблон и записываем результат в стандартный вывод
			if err := index.Execute(&buf, nil); err != nil {
				panic(err)
			}

			result := buf.String()

			g := General.Html{
				Html: result,
			}
			err = g.Save()
			if err != nil {
				fmt.Print(err)
			}
		}
	}
}
