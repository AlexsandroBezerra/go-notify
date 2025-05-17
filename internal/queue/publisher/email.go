package publisher

import (
	"AlexsandroBezerra/go-notify/internal/application/dto/message"
	"AlexsandroBezerra/go-notify/internal/queue/subject"
	"encoding/json"
	"github.com/nats-io/nats.go"
)

type EmailPublisher struct {
	natsConnection *nats.Conn
}

func NewEmailPublisher(natsConnection *nats.Conn) EmailPublisher {
	return EmailPublisher{natsConnection}
}

func (ep EmailPublisher) Publish(msg message.Email) (err error) {
	msgEncoded, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = ep.natsConnection.Publish(subject.Email, msgEncoded)
	return err
}
