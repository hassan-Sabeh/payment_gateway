// Package helpers provides helper structs and functions for handling API responses and generic functions.
package helpers

import (
	"math/rand"
	"time"
)

// GenerateRandomNumber generates a random number between 0 and 99 (inclusive).
func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

// IsEven checks if the given number is even.
// It returns true if the number is even, false otherwise.
func IsEven(number int) bool {
	return number%2 == 0
}