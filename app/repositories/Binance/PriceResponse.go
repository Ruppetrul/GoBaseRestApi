package Binance

type PriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
