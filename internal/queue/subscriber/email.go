package subscriber

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/message"
	"AlexsandroBezerra/go-notify/internal/application/usecase"
	repository "AlexsandroBezerra/go-notify/internal/storage/postgres"
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
	"log"
)

type EmailHandler struct {
	pgPool *pgxpool.Pool
}

func NewEmailHandler(pgPool *pgxpool.Pool) *EmailHandler {
	return &EmailHandler{pgPool}
}

func (eh *EmailHandler) ProcessMessage(msg *nats.Msg) {
	ctx := context.Background()

	var emailMsg message.Email
	err := json.Unmarshal(msg.Data, &emailMsg)
	if err != nil {
		log.Fatalln("Error unmarshalling emailMsg")
		return
	}

	processEmailUseCase := usecase.NewProcessEmail(eh.pgPool)
	err = processEmailUseCase.Execute(ctx, emailMsg)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("Email %s processed\n", emailMsg.ID)

	updateStatusUseCase := usecase.NewUpdateEmailStatus(eh.pgPool)
	err = updateStatusUseCase.Execute(ctx, emailMsg.ID, repository.DeliveryStatusDelivered)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = msg.Ack()
	if err != nil {
		log.Fatalln(err)
	}
}
