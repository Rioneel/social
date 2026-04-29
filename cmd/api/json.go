package main

import (
	"net/http"
	"encoding/json"
)
func writeJSON(w http.ResponseWriter, status int,data any)error{
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}