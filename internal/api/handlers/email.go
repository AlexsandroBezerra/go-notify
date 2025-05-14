package handlers

import (
	"AlexsandroBezerra/go-notify/internal/application/dto"
	"AlexsandroBezerra/go-notify/internal/application/usecase"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
)

type EmailHandler struct {
	databaseConnection *pgx.Conn
}

func NewEmailHandler(databaseConnection *pgx.Conn) *EmailHandler {
	return &EmailHandler{databaseConnection}
}

func (e *EmailHandler) ListEmails(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Route not implemented yet"))
}

func (e *EmailHandler) CreateEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var params dto.CreateEmailRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	useCase := usecase.NewCreateEmail(e.databaseConnection)
	id, err := useCase.Execute(ctx, params)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := dto.CreateEmailResponse{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json; utf-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
