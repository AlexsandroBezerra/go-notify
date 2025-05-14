package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeString(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(message))
	if err != nil {
		panic(err)
	}
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	fmt.Println(err)
	writeString(w, statusCode, err.Error())
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; utf-8")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(statusCode)
}
