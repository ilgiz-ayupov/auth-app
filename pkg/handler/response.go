package handler

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

type messageResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func sendResponse(w http.ResponseWriter, response any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}

func sendError(w http.ResponseWriter, err string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	error := errorResponse{Error: err, StatusCode: statusCode}
	json.NewEncoder(w).Encode(error)
}

func sendMessage(w http.ResponseWriter, msg string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	message := messageResponse{Message: msg, StatusCode: statusCode}
	json.NewEncoder(w).Encode(message)
}
