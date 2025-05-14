package handlers

import (
	"AlexsandroBezerra/go-notify/internal/application/usecase"
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
	useCase := usecase.NewCreateEmail(e.databaseConnection)
	useCase.Execute(ctx)
	w.WriteHeader(http.StatusCreated)
}
