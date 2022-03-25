package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(rw http.ResponseWriter, p interface{}, status int) {
	changeToByte, err := json.Marshal(p)
	if err != nil {
		http.Error(rw, "Error", http.StatusBadRequest)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write([]byte(changeToByte))
}
