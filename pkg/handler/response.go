package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type errorJSONResponse struct {
	Error string `json:"error"`
}

func sendResponseTemplate(w http.ResponseWriter, templatePath string) {
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := template.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func sendResponseJSON(w http.ResponseWriter, response any, statusCode int) {
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}

func sendResponseHTTP(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)

	w.Write([]byte(message))
}

func sendErrorJSON(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)

	error := errorJSONResponse{Error: err.Error()}
	json.NewEncoder(w).Encode(error)
}
