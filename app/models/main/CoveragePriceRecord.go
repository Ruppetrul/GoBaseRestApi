package main

type CoveragePriceRecord struct {
	RowDate       string  `json:"rowDate"`
	LastMaxRaw    string  `json:"last_maxRaw"`
	LastMinRaw    string  `json:"last_minRaw"`
	RowDateRaw    int64   `json:"rowDateRaw"`
	LastCloseRaw  string  `json:"last_closeRaw"`
	LastOpenRaw   string  `json:"last_openRaw"`
	ChangePercent float64 `json:"change_precentRaw"`
}
