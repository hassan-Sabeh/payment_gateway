// Package models provides data models used in the application.
package models

import (
	"time"
)

// Payment represents a payment.
type Payment struct {
	ID          string    `json:"payment_id"`
	CardInfo    Card      `json:"card_info"`
	ProcessedAt time.Time `json:"processed_at"`
	Status      string    `json:"payment_status"`
	// User info or id should be included here.
}