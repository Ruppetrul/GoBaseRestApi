package workers

import (
	"fmt"
	"time"
)

func RegisterCurrentPriceWorker() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			//TODO logic
			fmt.Println("Current time: ", t)
		}
	}
}
