package usecase

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/response"
	repository "AlexsandroBezerra/go-notify/internal/storage/postgres"
	"context"
	"github.com/jackc/pgx/v5"
)

type ListEmail struct {
	databaseConnection *pgx.Conn
}

func NewListEmail(databaseConnection *pgx.Conn) *ListEmail {
	return &ListEmail{databaseConnection}
}

func (l *ListEmail) Execute(ctx context.Context) (response.ListEmail, error) {
	queries := repository.New(l.databaseConnection)

	emails, err := queries.ListEmails(ctx)
	if err != nil {
		return response.ListEmail{}, err
	}

	result := response.ListEmail{}
	for _, email := range emails {
		result = append(result, response.Email{
			ID:        email.ID.String(),
			Recipient: email.Recipient,
			Subject:   email.Subject,
			Body:      email.Body,
			Priority:  email.Priority,
			Status:    string(email.Status),
			CreatedAt: email.CreatedAt.Time,
		})
	}

	return result, nil
}
