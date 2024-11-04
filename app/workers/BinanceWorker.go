package workers

import (
	"firstRest/repositories"
	"fmt"
	"time"
)

func RegisterCoinGeckoWorker() {
	ticker := time.NewTicker(5 * time.Second)
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
				err, _ := ticker.Save()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
