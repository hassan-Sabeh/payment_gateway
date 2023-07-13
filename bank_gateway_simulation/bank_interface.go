// Package BankGatewaySimulation provides functionality related to the bank gateway simulation.
package BankGatewaySimulation

import (
	"time"

	"github.com/google/uuid"

	m "github.com/hassan-Sabeh/payment_gateway/models"
)

// BankResponse represents the response received from the bank.
type BankResponse struct {
	Status    string    `json:"response_status"`
	UID       uuid.UUID `json:"transaction_id"`
	TimeStamp time.Time `json:"created_at"`
}

// BankInterface defines the interface for interacting with the bank.
type BankInterface interface {
	PaymentTransaction(*m.Card) BankResponse
}