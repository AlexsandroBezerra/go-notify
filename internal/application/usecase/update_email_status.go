package usecase

import (
	repository "AlexsandroBezerra/go-notify/internal/storage/postgres"
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UpdateEmailStatus struct {
	pgPool *pgxpool.Pool
}

func NewUpdateEmailStatus(pgPool *pgxpool.Pool) *UpdateEmailStatus {
	return &UpdateEmailStatus{pgPool}
}

func (c *UpdateEmailStatus) Execute(ctx context.Context, id string, status repository.DeliveryStatus) (err error) {
	queries := repository.New(c.pgPool)

	emailId := pgtype.UUID{}
	err = emailId.Scan(id)
	if err != nil {
		return err
	}

	return queries.CreateEmailStatus(ctx, repository.CreateEmailStatusParams{
		EmailID: emailId,
		Status:  status,
	})
}
