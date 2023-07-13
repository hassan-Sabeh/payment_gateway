// Package handlers provides request handlers for processing payments and retrieving payment information.
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"

	bank "github.com/hassan-Sabeh/payment_gateway/bank_gateway_simulation"
	db "github.com/hassan-Sabeh/payment_gateway/database"
	h "github.com/hassan-Sabeh/payment_gateway/helpers"
	m "github.com/hassan-Sabeh/payment_gateway/models"
)

// ProcessPayment handles the request for processing a payment.
// It receives a card object, initiates the payment with the bank gateway,
// saves the bank response in the database, and returns the processed payment.
func ProcessPayment(c echo.Context) error {
	card := new(m.Card)
	// Check for validation when binding the credit card object.
	if err := c.Bind(card); err != nil {
		// TODO: Add validator for providing feedback response to the client on which field is invalid.
		return c.JSON(http.StatusUnprocessableEntity, h.NewErrorResponse("Invalid credit card"))
	}

	// Initiate payment with the bank gateway.
	bank := c.Get("bank").(*bank.BankGateway)
	bankResponse := bank.PaymentTransaction(card)

	// Mask the credit card number.
	maskedCard := card.MaskCreditCardNumber()

	// Create the payment object with the bank response.
	currentPayment := m.Payment{
		ID:          bankResponse.UID.String(),
		CardInfo:    maskedCard,
		ProcessedAt: bankResponse.TimeStamp,
		Status:      bankResponse.Status,
	}

	// Save the bank response in the database for traceability.
	db := c.Get("db").(*db.Database)
	paymentFromDb, err := db.InsertPayment(currentPayment)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, h.NewErrorResponse("Error persisting payment"))
	}

	return c.JSON(http.StatusAccepted, h.NewSuccessResponse(paymentFromDb))
}

// GetPayment handles the request for retrieving payment information.
// It receives the payment ID, retrieves the payment from the database,
// and returns the payment information.
func GetPayment(c echo.Context) error {
	id := c.Param("id")

	// Retrieve the database object from the context.
	db := c.Get("db").(*db.Database)
	payment, err := db.SelectPayment(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, h.NewErrorResponse("Payment not found"))
	}

	return c.JSON(http.StatusOK, h.NewSuccessResponse(*payment))
}