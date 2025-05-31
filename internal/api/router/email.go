package router

import (
	"AlexsandroBezerra/go-notify/internal/api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

type EmailRouter struct {
	dbPool         *pgxpool.Pool
	natsConnection *nats.Conn
}

func NewEmailRouter(dbPool *pgxpool.Pool, natsConnection *nats.Conn) *EmailRouter {
	return &EmailRouter{dbPool, natsConnection}
}

func (e *EmailRouter) RegisterRoutes(r chi.Router) {
	emailHandler := handler.NewEmailHandler(e.dbPool, e.natsConnection)

	r.Route("/emails", func(r chi.Router) {
		r.Get("/", emailHandler.ListEmails)
		r.Post("/", emailHandler.CreateEmail)
	})
}
