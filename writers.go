package main

import (
	"encoding/json"
	"net/http"
)


func writeJSON(w http.ResponseWriter, value any) error {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	return encoder.Encode(value)	
}

