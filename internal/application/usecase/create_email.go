package usecase

import (
	repository "AlexsandroBezerra/go-notify/internal/storage/postgres"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type CreateEmail struct {
	queries *repository.Queries
}

func NewCreateEmail(databaseConnection *pgx.Conn) *CreateEmail {
	queries := repository.New(databaseConnection)
	return &CreateEmail{queries}
}

func (c *CreateEmail) Execute(ctx context.Context) {
	email, err := c.queries.CreateEmail(ctx, repository.CreateEmailParams{
		Recipient: "me@alexsandrobezerra.dev",
		Subject:   "Test",
		Body:      "Just a message...",
		Priority:  1,
		Status:    repository.DeliveryStatus("pending"),
	})
	if err != nil {
		panic(err) // TODO: handle error
	}
	fmt.Println(email)
}
