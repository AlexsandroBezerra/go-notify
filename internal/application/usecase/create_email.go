package usecase

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/request"
	repository "AlexsandroBezerra/go-notify/internal/storage/postgres"
	"context"
	"github.com/jackc/pgx/v5"
)

type CreateEmail struct {
	queries *repository.Queries
}

func NewCreateEmail(databaseConnection *pgx.Conn) *CreateEmail {
	queries := repository.New(databaseConnection)
	return &CreateEmail{queries}
}

func (c *CreateEmail) Execute(ctx context.Context, params request.CreateEmail) (string, error) {
	email, err := c.queries.CreateEmail(ctx, repository.CreateEmailParams{
		Recipient: params.Recipient,
		Subject:   params.Subject,
		Body:      params.Body,
		Priority:  params.Priority,
		Status:    repository.DeliveryStatusPending,
	})
	if err != nil {
		return "", err
	}

	return email.String(), nil
}
