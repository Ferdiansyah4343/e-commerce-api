package helpers

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, message, status string, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := map[string]interface{}{
		"message": message,
		"status":  status,
		"data":    data,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response to JSON", http.StatusInternalServerError)
	}
}
