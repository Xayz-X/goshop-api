package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func WriterJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriterError(w http.ResponseWriter, status int, erre error) error {
	return WriterJSON(w, status, map[string]string{"error": erre.Error()})
}

func WriteValidationError(w http.ResponseWriter, status int, err error) error {

	if vErrs, ok := err.(validator.ValidationErrors); ok {
		validationErrors := make(map[string]string)
		for _, vErr := range vErrs {
			validationErrors[vErr.Field()] = vErr.Tag()
		}
		return WriterJSON(w, status, map[string]interface{}{"error": validationErrors})
	}
	return WriterJSON(w, status, map[string]string{"error": err.Error()})
}
