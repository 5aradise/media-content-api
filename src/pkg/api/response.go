package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteHTMLf(w http.ResponseWriter, statusCode int, formatHTML string, a ...any) {
	const op = "api.wtireHTML"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(fmt.Sprintf(formatHTML, a...)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("%s: %v\n", op, err)
	}
}

func WriteTextf(w http.ResponseWriter, statusCode int, format string, a ...any) {
	const op = "api.wtireTextf"

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(fmt.Sprintf(format, a...)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("%s: %v\n", op, err)
	}
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) {
	const op = "api.wtireJSON"

	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%s: %v\n", op, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("%s: %v\n", op, err)
	}
}

func WriteErrorf(w http.ResponseWriter, statusCode int, format string, a ...any) {
	const op = "api.wtireErrorf"

	errMsg := fmt.Sprintf(format, a...)
	if statusCode >= 500 {
		log.Printf("%s: %s\n", op, errMsg)
		errMsg = "internal server error"
	}

	data, err := json.Marshal(ErrorResponse{errMsg})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("%s: %v\n", op, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("%s: %v\n", op, err)
	}
}

func WriteNoContent(w http.ResponseWriter) {
	const op = "api.wtireNoContent"

	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("%s: %v\n", op, err)
	}
}
