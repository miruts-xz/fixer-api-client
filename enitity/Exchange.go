package enitity

import "time"

type Currency string

// Constant representing types of api requests

type SymbolsRequest struct {
	Success bool                `json:"success"`
	Symbols map[Currency]string `json:"symbols"`
}

// Rate defines a single exchange rate
type Rate map[Currency]float64

// Exchange struct represents a currency exchange rate
type ExchangeRequest struct {
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
	Date      time.Time `json:"date"`
	Base      Currency  `json:"base"`
	Rates     Rate      `json:"rates"`
}
