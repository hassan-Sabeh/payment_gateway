// Package helpers provides helper structs and functions for handling API responses and generic functions.
package helpers

// ErrorResponse represents an error response returned by the API.
type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse creates a new ErrorResponse with the given error message.
func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Message: message,
	}
}

// SuccessResponse represents a success response returned by the API.
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// NewSuccessResponse creates a new SuccessResponse with the given data.
func NewSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Data: data,
	}
}