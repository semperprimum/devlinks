package utils

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidationErrorResponse represents a JSON response for validation errors
type ValidationErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidateStruct(data interface{}) []*ValidationErrorResponse {
	var errors []*ValidationErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func WriteValidationError(w http.ResponseWriter, validationErrors []*ValidationErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(validationErrors)
}
