package router

import (
	"AlexsandroBezerra/go-notify/internal/api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/nats.go"
)

type EmailRouter struct {
	databaseConnection *pgx.Conn
	natsConnection     *nats.Conn
}

func NewEmailRouter(databaseConnection *pgx.Conn, natsConnection *nats.Conn) *EmailRouter {
	return &EmailRouter{databaseConnection, natsConnection}
}

func (e *EmailRouter) RegisterRoutes(r chi.Router) {
	emailHandler := handler.NewEmailHandler(e.databaseConnection, e.natsConnection)

	r.Route("/emails", func(r chi.Router) {
		r.Get("/", emailHandler.ListEmails)
		r.Post("/", emailHandler.CreateEmail)
	})
}
