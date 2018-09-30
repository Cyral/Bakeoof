package main

import (
	"encoding/json"
	"net/http"
)

type okResponse struct {
	Ok bool `json:"ok"`
  Data interface{} `json:"data"`
}

type errResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

// RootController handles the / endpoint
func RootController(w http.ResponseWriter, r *http.Request) {
	sendResponse("Hello, World!", w)
}

func sendResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(okResponse{Ok: true, Data: data})
}

func sendError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(errResponse{Ok: false, Error: err.Error()})
}
