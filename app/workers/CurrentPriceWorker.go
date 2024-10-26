package workers

import (
	"firstRest/repositories"
	"fmt"
	"time"
)

func RegisterCurrentPriceWorker() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			binanceRep := repositories.BinanceRepository{}
			prices := binanceRep.GetCurrentPrices()

			if prices == nil {
				fmt.Println("Prices is empty")
			}
			fmt.Println("Current time: ", prices)
		}
	}
}
