package usecase

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/response"
	repository "AlexsandroBezerra/go-notify/internal/storage/postgres"
	"context"
	"github.com/jackc/pgx/v5"
)

type ListEmail struct {
	queries *repository.Queries
}

func NewListEmail(databaseConnection *pgx.Conn) *ListEmail {
	queries := repository.New(databaseConnection)
	return &ListEmail{queries}
}

func (l *ListEmail) Execute(ctx context.Context) (response.ListEmail, error) {
	emails, err := l.queries.ListEmails(ctx)
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
