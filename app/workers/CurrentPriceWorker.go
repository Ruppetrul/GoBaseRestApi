package workers

import (
	"firstRest/repositories/Binance"
	"fmt"
	"time"
)

func RegisterCurrentPriceWorker() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			binanceRep := Binance.Repository{}
			prices, err := binanceRep.GetCurrentPrices()

			if err != nil {
				fmt.Println(err)
			}
			if prices == nil {
				fmt.Println("Prices is empty")
			}

			for _, price := range prices {
				if len(price.Name) > 3 && price.Name[len(price.Name)-3:] == "USD" {
					err := price.Save()
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}
