package usecase

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/response"
	repository "AlexsandroBezerra/go-notify/internal/storage/postgres"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ListEmail struct {
	dbPool *pgxpool.Pool
}

func NewListEmail(dbPool *pgxpool.Pool) *ListEmail {
	return &ListEmail{dbPool}
}

func (l *ListEmail) Execute(ctx context.Context) (response.ListEmail, error) {
	queries := repository.New(l.dbPool)

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
			CreatedAt: email.CreatedAt.Time,
		})
	}

	return result, nil
}
