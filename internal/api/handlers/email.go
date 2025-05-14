package handlers

import (
	"AlexsandroBezerra/go-notify/internal/application/usecase"
	"net/http"
)

func ListEmails(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Route not implemented yet"))
}

func CreateEmail(w http.ResponseWriter, r *http.Request) {
	useCase := usecase.NewCreateEmail()
	useCase.Execute()
	w.WriteHeader(http.StatusCreated)
}
