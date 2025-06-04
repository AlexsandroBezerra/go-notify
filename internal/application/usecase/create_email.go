package usecase

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/message"
	"AlexsandroBezerra/go-notify/internal/application/dto/request"
	"AlexsandroBezerra/go-notify/internal/application/model"
	"AlexsandroBezerra/go-notify/internal/queue/publisher"
	repository "AlexsandroBezerra/go-notify/internal/storage/postgres"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

type CreateEmail struct {
	pgPool         *pgxpool.Pool
	natsConnection *nats.Conn
}

func NewCreateEmail(pgPool *pgxpool.Pool, natsConnection *nats.Conn) *CreateEmail {
	return &CreateEmail{pgPool, natsConnection}
}

func (c *CreateEmail) Execute(ctx context.Context, params request.CreateEmail) (string, error) {
	queries := repository.New(c.pgPool)
	trx, err := c.pgPool.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer trx.Rollback(ctx)

	queriesTrx := queries.WithTx(trx)

	emailId, err := model.NewId()
	if err != nil {
		return "", err
	}

	err = queriesTrx.CreateEmail(ctx, repository.CreateEmailParams{
		ID:        emailId.PgId(),
		Recipient: params.Recipient,
		Subject:   params.Subject,
		Body:      params.Body,
		Priority:  params.Priority,
	})
	if err != nil {
		return "", err
	}

	emailStatusId, err := model.NewId()
	if err != nil {
		return "", err
	}

	err = queriesTrx.CreateEmailStatus(ctx, repository.CreateEmailStatusParams{
		ID:      emailStatusId.PgId(),
		EmailID: emailId.PgId(),
		Status:  repository.DeliveryStatusPending,
	})
	if err != nil {
		return "", err
	}

	emailPublisher := publisher.NewEmailPublisher(c.natsConnection)

	err = emailPublisher.Publish(message.Email{
		ID:        emailId.String(),
		Recipient: params.Recipient,
		Subject:   params.Subject,
		Body:      params.Body,
		Priority:  params.Priority,
	})
	if err != nil {
		return "", err
	}

	err = trx.Commit(ctx)
	if err != nil {
		return "", err
	}

	return emailId.String(), nil
}
