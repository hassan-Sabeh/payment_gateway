// Package database provides functionality for simulating a database.
package database

import (
	"errors"

	m "github.com/hassan-Sabeh/payment_gateway/models"
)

// Database represents the simulated database.
type Database struct {
	entries map[string]*m.Payment
}

// NewDbConnection creates a new database connection and returns a pointer to the Database struct.
func NewDbConnection() *Database {
	return new(Database)
}

// SelectPayment retrieves a payment from the database based on the given ID.
// It returns the payment and an error if the payment is not found.
func (db *Database) SelectPayment(id string) (payment *m.Payment, err error) {
	// Add more validators to id to avoid useless calls to the database.
	payment, ok := db.entries[id]
	if !ok {
		return nil, errors.New("Payment not found")
	}
	return payment, nil
}

// InsertPayment inserts a payment into the database.
// It returns the inserted payment and an error if the payment already exists.
func (db *Database) InsertPayment(p m.Payment) (payment *m.Payment, err error) {
	// Here ID is duplicated used as a primary key and is a field in payment (only for this test).
	_, ok := db.entries[p.ID]
	if ok {
		return nil, errors.New("Payment already exists")
	}
	if db.entries == nil {
		db.entries = make(map[string]*m.Payment)
	}
	db.entries[p.ID] = &p
	return &p, nil
}