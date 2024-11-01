package api

import (
	"encoding/json"
	"net/http"
)

// Coin bal
type CoinBalanceParams  struct {
	Username string
}

// Coin Bal response
type CoinBalanceResponse struct {
	Code int // Status Code 
	Balance int64 // Balance
}

type Error struct {
	Code int // Error code
	Message string // Error message
}

func writeError (w http.ResponseWriter, message string, code int){
	resp := Error {
		Code: code, 
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpected Error Occured", http.StatusInternalServerError)
	}
)