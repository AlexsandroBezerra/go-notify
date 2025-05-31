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
	workerId           int
	postgresConnection *pgx.Conn
}

func NewEmailHandler(WorkerId int, postgresConnection *pgx.Conn) *EmailHandler {
	return &EmailHandler{WorkerId, postgresConnection}
}

func (eh *EmailHandler) ProcessMessage(msg *nats.Msg) {
	ctx := context.Background()

	var emailMsg message.Email
	err := json.Unmarshal(msg.Data, &emailMsg)
	if err != nil {
		log.Fatalf("[Worker %d] Error unmarshalling emailMsg\n", eh.workerId)
		return
	}

	// TODO: Send email in usecase

	log.Printf("[Worker %d] Updating status to delivered to emailId: %s\n", eh.workerId, emailMsg.ID)

	err = eh.updateStatus(ctx, emailMsg.ID, repository.DeliveryStatusDelivered)
	if err != nil {
		log.Fatalf("[Worker %d] %s\n", eh.workerId, err)
	}
}

// TODO: Move to usecase
func (eh *EmailHandler) updateStatus(ctx context.Context, ID string, status repository.DeliveryStatus) (err error) {
	queries := repository.New(eh.postgresConnection)
	emailId := pgtype.UUID{}
	err = emailId.Scan(ID)
	if err != nil {
		return errors.New("error scanning message uuid to update status")
	}

	err = queries.CreateEmailStatus(ctx, repository.CreateEmailStatusParams{
		EmailID: emailId,
		Status:  status,
	})
	if err != nil {
		return errors.New("error updating email status")
	}

	return nil
}
