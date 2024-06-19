package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a JSON error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// WriteJSONError writes an error message as a JSON response
func WriteJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
