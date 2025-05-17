package subscriber

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/message"
	repository "AlexsandroBezerra/go-notify/internal/storage/postgres"
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nats-io/nats.go"
	"log"
)

type EmailHandler struct {
	workerId int
	queries  *repository.Queries
}

func NewEmailHandler(WorkerId int, postgresConnection *pgx.Conn) *EmailHandler {
	queries := repository.New(postgresConnection)
	return &EmailHandler{WorkerId, queries}
}

func (eh *EmailHandler) ProcessMessage(msg *nats.Msg) {
	ctx := context.Background()

	var emailMsg message.Email
	err := json.Unmarshal(msg.Data, &emailMsg)
	if err != nil {
		log.Fatalf("[Worker %d] Error unmarshalling emailMsg\n", eh.workerId)
		return
	}

	// TODO: Send email

	log.Printf("[Worker %d] Updating status to delivered to emailId: %s\n", eh.workerId, emailMsg.ID)

	err = eh.updateStatus(ctx, emailMsg.ID, repository.DeliveryStatusDelivered)
	if err != nil {
		log.Fatalf("[Worker %d] %s\n", eh.workerId, err)
	}
}

func (eh *EmailHandler) updateStatus(ctx context.Context, ID string, status repository.DeliveryStatus) (err error) {
	emailId := pgtype.UUID{}
	err = emailId.Scan(ID)
	if err != nil {
		return errors.New("error scanning message uuid to update status")
	}

	err = eh.queries.UpdateEmailStatus(ctx, repository.UpdateEmailStatusParams{
		ID:     emailId,
		Status: status,
	})
	if err != nil {
		return errors.New("error updating email status")
	}

	return nil
}
