// Package BankGatewaySimulation provides functionality related to the bank gateway simulation.
package BankGatewaySimulation

import (
	"time"

	"github.com/google/uuid"

	help "github.com/hassan-Sabeh/payment_gateway/helpers"
	m "github.com/hassan-Sabeh/payment_gateway/models"
)

// BankGateway represents the implementation of the BankInterface.
type BankGateway struct{}

// PaymentTransaction simulates a payment transaction with the bank.
// It receives a Card object and returns a BankResponse.
func (b *BankGateway) PaymentTransaction(card *m.Card) BankResponse {
	timeStamp := time.Now()
	uid := uuid.New()
	if paymentSuccess := paymentSuccess(); paymentSuccess {
		// Payment successful
		return BankResponse{
			Status:    "Success",
			UID:       uid,
			TimeStamp: timeStamp,
		}
	}
	// Payment failed for whatever reason
	return BankResponse{
		Status:    "Fail",
		UID:       uid,
		TimeStamp: timeStamp,
	}
}

// paymentSuccess generates a random success or failure for simulation purposes.
func paymentSuccess() bool {
	randNum := help.GenerateRandomNumber()
	return help.IsEven(randNum)
}