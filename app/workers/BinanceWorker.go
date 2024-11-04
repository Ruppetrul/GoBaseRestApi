package workers

import (
	"firstRest/repositories"
	"fmt"
	"time"
)

func RegisterCoinGeckoWorker() {
	ticker := time.NewTicker(6 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tickers, err := repositories.CoinGeckoRepository{}.GetTicker()

			if err != nil {
				fmt.Println(err)
			}
			if tickers == nil {
				fmt.Println("Prices is empty")
			}

			for _, ticker := range tickers {
				_, err := ticker.Save()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
