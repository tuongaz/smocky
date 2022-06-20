package server

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Error string `json:"error"`
}

func responseError(w http.ResponseWriter, status int, message string) {
	resp := APIError{message}
	response(w, status, resp)
}

func response(w http.ResponseWriter, status int, data any) {
	body, _ := json.Marshal(data)
	w.WriteHeader(status)
	_, _ = w.Write(body)
}
