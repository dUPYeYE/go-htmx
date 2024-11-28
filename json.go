package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithError responds with an error message and status code.
// If the status code is greater than 499, it logs the error message.
//
// Parameters:
// - w: http.ResponseWriter - Response writer
// - code: int - Status code
// - msg: string - Error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

// respondWithError responds with an error message and status code.
// If the status code is greater than 499, it logs the error message.
//
// Parameters:
// - w: http.ResponseWriter - Response writer
// - code: int - Status code
// - payload: interface{} - Payload to be marshalled to JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	if _, err = w.Write(dat); err != nil {
		log.Printf("Error writing response: %s", err)
	}
}
