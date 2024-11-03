package workers

import (
	"firstRest/models"
	"firstRest/models/General"
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
			html := "<html lang=\"en\">\n    <table>\n        <tr>\n            <th>\n                Name\n            </th>\n            <th>\n                Price\n            </th>\n            <th>\n                24h Volume\n            </th>\n        </tr>\n        <tr>\n            <td>\n                Test\n            </td>\n            <td>\n                123.0\n            </td>\n            <td>\n                77777\n            </td>\n        </tr>\n        <tr>\n            <td>\n                Test\n            </td>\n            <td>\n                123.0\n            </td>\n            <td>\n                77777\n            </td>\n        </tr>\n        <tr>\n            <td>\n                Test\n            </td>\n            <td>\n                123.0\n            </td>\n            <td>\n                77777\n            </td>\n        </tr>\n        <tr>\n            <td>\n                Test\n            </td>\n            <td>\n                123.0\n            </td>\n            <td>\n                77777\n            </td>\n        </tr>\n        <tr>\n            <td>\n                Test\n            </td>\n            <td>\n                123.0\n            </td>\n            <td>\n                77777\n            </td>\n        </tr>\n        <tr>\n            <td>\n                Test\n            </td>\n            <td>\n                123.0\n            </td>\n            <td>\n                77777\n            </td>\n        </tr>\n        <tr>\n            <td>\n                Test\n            </td>\n            <td>\n                123.0\n            </td>\n            <td>\n                77777\n            </td>\n        </tr>\n    </table>\n</html>"
			g := General.Html{
				Html: html,
			}
			err = g.Save()
			if err != nil {
				fmt.Print(err)
			}
		}
	}
}
