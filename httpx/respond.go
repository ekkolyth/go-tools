package httpx

import (
	"encoding/json"
	"net/http"
)

// write a JSON response with status code
func RespondJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// write error json response
func RespondError(w http.ResponseWriter, status int, msg string) {
	RespondJSON(w, status, map[string]any{
		"error":   http.StatusText(status),
		"message": msg,
	})
}
