package model

type Stock struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}