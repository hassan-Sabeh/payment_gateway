// Package models provides data models used in the application.
package models

// Card represents a credit card.
type Card struct {
	Number      string  `json:"card_number"`
	ExpiryMonth int     `json:"expiry_month"`
	ExpiryYear  int     `json:"expiry_year"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	CVV         string  `json:"cvv"`
}

// MaskCreditCardNumber masks the credit card number of the Card.
// It masks all digits except the last four.
// It returns a new Card object with the masked credit card number.
func (c *Card) MaskCreditCardNumber() Card {
	maskedNumber := maskCreditCard(c.Number)
	c.Number = maskedNumber
	return *c
}

// maskCreditCard masks the given credit card number.
func maskCreditCard(number string) string {
	length := len(number)
	masked := "XXXXXXXXXXXX"
	masked += number[length-4:]
	return masked
}
