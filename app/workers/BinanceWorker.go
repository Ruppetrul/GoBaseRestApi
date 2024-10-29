package workers

import (
	"firstRest/models/binance"
	"firstRest/repositories/Binance"
	"fmt"
	"time"
)

func RegisterBinanceWorker() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			binanceRep := Binance.Repository{}
			tickers, err := binanceRep.GetTicker()

			if err != nil {
				fmt.Println(err)
			}
			if tickers == nil {
				fmt.Println("Prices is empty")
			}

			for _, ticker := range tickers {
				if len(ticker.Symbol) > 3 && ticker.Symbol[len(ticker.Symbol)-3:] == "USD" {
					b := binance.Ticker{
						Symbol: ticker.Symbol, LastPrice: ticker.LastPrice,
						PriceChangePercent: ticker.PriceChangePercent, QuoteVolume: ticker.QuoteVolume,
					}
					err := b.Save()
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}
