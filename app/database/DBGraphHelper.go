package database

import (
	"fmt"
	"time"
)

type CoveragePriceRecord struct {
	Pair         string
	Date         string
	AveragePrice float64
}

type HalvingRecord struct {
	Date time.Time
}

func GetAverageData(pair string, halving time.Time, lastHalving time.Time) ([]CoveragePriceRecord, error) {
	rows, err := Select(fmt.Sprintf("SELECT date, average_price"+
		" FROM average_price_per_day "+
		"WHERE pair = '%s' "+
		"AND date >= '%s' AND date < '%s'"+
		"ORDER BY date DESC;", pair,
		lastHalving.Format("2006-01-02"), halving.Format("2006-01-02")))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []CoveragePriceRecord

	for rows.Next() {
		var record CoveragePriceRecord
		if err := rows.Scan(
			&record.Date, &record.AveragePrice,
		); err != nil {
			return nil, err
		}
		record.Pair = pair
		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func GetHalvings() ([]HalvingRecord, error) {
	rows, err := Select(fmt.Sprintf("SELECT halving_date as date FROM bitcoin_halving_dates;"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var records []HalvingRecord
	for rows.Next() {
		var record HalvingRecord
		if err := rows.Scan(&record.Date); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return records, nil
}
