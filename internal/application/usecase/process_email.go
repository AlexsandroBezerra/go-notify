package usecase

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/message"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type ProcessEmail struct {
	pgPool *pgxpool.Pool
}

func NewProcessEmail(pgPool *pgxpool.Pool) *ProcessEmail {
	return &ProcessEmail{pgPool}
}

func (c *ProcessEmail) Execute(ctx context.Context, email message.Email) (err error) {
	log.Println("Processing email: " + email.ID)
	return nil
}
