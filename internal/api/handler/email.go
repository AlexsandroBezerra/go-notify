package handler

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/request"
	"AlexsandroBezerra/go-notify/internal/application/dto/response"
	"AlexsandroBezerra/go-notify/internal/application/usecase"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/nats.go"
	"net/http"
)

type EmailHandler struct {
	databaseConnection *pgx.Conn
	natsConnection     *nats.Conn
}

func NewEmailHandler(databaseConnection *pgx.Conn, natsConnection *nats.Conn) *EmailHandler {
	return &EmailHandler{databaseConnection, natsConnection}
}

func (e *EmailHandler) ListEmails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	useCase := usecase.NewListEmail(e.databaseConnection)
	emails, err := useCase.Execute(ctx)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, http.StatusOK, emails)
}

func (e *EmailHandler) CreateEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var params request.CreateEmail
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	useCase := usecase.NewCreateEmail(e.databaseConnection, e.natsConnection)
	id, err := useCase.Execute(ctx, params)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusCreated, response.CreateEmail{
		ID: id,
	})
}
