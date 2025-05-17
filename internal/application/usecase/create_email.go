package usecase

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/message"
	"AlexsandroBezerra/go-notify/internal/application/dto/request"
	"AlexsandroBezerra/go-notify/internal/queue/publisher"
	repository "AlexsandroBezerra/go-notify/internal/storage/postgres"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/nats.go"
)

type CreateEmail struct {
	queries        *repository.Queries
	emailPublisher publisher.EmailPublisher
}

func NewCreateEmail(databaseConnection *pgx.Conn, natsConnection *nats.Conn) *CreateEmail {
	queries := repository.New(databaseConnection)
	emailPublisher := publisher.NewEmailPublisher(natsConnection)
	return &CreateEmail{queries, emailPublisher}
}

func (c *CreateEmail) Execute(ctx context.Context, params request.CreateEmail) (string, error) {
	emailId, err := c.queries.CreateEmail(ctx, repository.CreateEmailParams{
		Recipient: params.Recipient,
		Subject:   params.Subject,
		Body:      params.Body,
		Priority:  params.Priority,
		Status:    repository.DeliveryStatusPending,
	})
	if err != nil {
		return "", err
	}
	err = c.emailPublisher.Publish(message.Email{
		ID:        emailId.String(),
		Recipient: params.Recipient,
		Subject:   params.Subject,
		Body:      params.Body,
		Priority:  params.Priority,
	})
	if err != nil {
		return "", err
	}

	return emailId.String(), nil
}
