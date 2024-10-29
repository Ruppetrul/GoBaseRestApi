package workers

import (
	"firstRest/models"
	"firstRest/models/binance"
	"fmt"
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
		}
	}
}
