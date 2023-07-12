package models

type Card struct{
	Number   string `json:"card_number"`
    ExpiryMonth  int    `json:"expiry_month"`
    ExpiryYear   int    `json:"expiry_year"`
    Amount       float64 `json:"amount"`
    Currency     string `json:"currency"`
    CVV          string `json:"cvv"`
}