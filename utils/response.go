package utils

import (
	"encoding/json"
	"net/http"
)

func WriterJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func WriterError(w http.ResponseWriter, status int, erre error) error {
	return WriterJSON(w, status, map[string]string{"error": erre.Error()})
}
