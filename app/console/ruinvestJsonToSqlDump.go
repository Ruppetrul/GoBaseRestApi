// Parse API response from ru invest and generate sql dump
package console

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type CoveragePriceRecord struct {
	RowDate       string  `json:"rowDate"`
	LastMaxRaw    string  `json:"last_maxRaw"`
	LastMinRaw    string  `json:"last_minRaw"`
	RowDateRaw    int64   `json:"rowDateRaw"`
	LastCloseRaw  string  `json:"last_closeRaw"`
	LastOpenRaw   string  `json:"last_openRaw"`
	ChangePercent float64 `json:"change_precentRaw"`
}

func main() {
	data, err := os.ReadFile("btc_import.json")
	if err != nil {
		log.Fatal("Ошибка при чтении файла: ", err)
	}

	var records []CoveragePriceRecord
	err = json.Unmarshal(data, &records)
	if err != nil {
		log.Fatal("Ошибка при разборе JSON: ", err)
	}

	file, err := os.Create("btc_day_history_dump.sql")
	if err != nil {
		log.Fatal("Ошибка при создании файла: ", err)
	}
	defer file.Close()

	fmt.Println("-- SQL дамп для вставки данных в таблицу average_price_per_day")
	for _, record := range records {
		parsedDate := time.Unix(record.RowDateRaw, 0)
		formattedDate := parsedDate.Format("2006-01-02")
		fmt.Printf("Data: %v \n", formattedDate)

		lastMaxRaw, err := strconv.ParseFloat(record.LastMaxRaw, 64)
		if err != nil {
			log.Fatal("Ошибка при преобразовании last_maxRaw: ", err)
		}

		lastMinRaw, err := strconv.ParseFloat(record.LastMinRaw, 64)
		if err != nil {
			log.Fatal("Ошибка при преобразовании last_minRaw: ", err)
		}

		averagePrice := (lastMaxRaw + lastMinRaw) / 2

		sql := fmt.Sprintf("INSERT INTO average_price_per_day (pair, date, average_price) VALUES ('BTC/USD', '%s', %.8f);\n", formattedDate, averagePrice)
		_, err = file.WriteString(sql)
		if err != nil {
			log.Fatal(err)
		}
	}
}
