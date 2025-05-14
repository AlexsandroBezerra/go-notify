package router

import (
	"AlexsandroBezerra/go-notify/internal/api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type EmailRouter struct {
	databaseConnection *pgx.Conn
}

func NewEmailRouter(databaseConnection *pgx.Conn) *EmailRouter {
	return &EmailRouter{databaseConnection}
}

func (e *EmailRouter) RegisterRoutes(r chi.Router) {
	emailHandler := handler.NewEmailHandler(e.databaseConnection)

	r.Route("/emails", func(r chi.Router) {
		r.Get("/", emailHandler.ListEmails)
		r.Post("/", emailHandler.CreateEmail)
	})
}
