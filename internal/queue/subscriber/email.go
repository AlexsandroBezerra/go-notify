package subscriber

import (
	"github.com/nats-io/nats.go"
	"log"
)

type EmailHandler struct {
	WorkerId int
}

func NewEmailHandler(WorkerId int) *EmailHandler {
	return &EmailHandler{WorkerId}
}

func (eh *EmailHandler) ProcessMessage(msg *nats.Msg) {
	log.Printf("[Worker %d (%d)] Received message: %s\n", eh.WorkerId, msg.Data)
}
