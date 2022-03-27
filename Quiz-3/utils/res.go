package util

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	encoded, err := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "Oops...", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(encoded))
}
