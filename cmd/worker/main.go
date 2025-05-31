package main

import (
	"AlexsandroBezerra/go-notify/internal/queue/subject"
	"AlexsandroBezerra/go-notify/internal/queue/subscriber"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

const WorkerCount = 5

func main() {
	ctx := context.Background()

	pgPool, err := pgxpool.New(ctx, "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	natsConnection, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer natsConnection.Close()

	wg := sync.WaitGroup{}

	worker := func(id int) {
		defer wg.Done()
		handler := subscriber.NewEmailHandler(pgPool)
		subscription, err := natsConnection.QueueSubscribe(subject.Email, "email-queue", handler.ProcessMessage)
		if err != nil {
			log.Printf("Error subscribing worker %d: %v\n", id, err)
			return
		}
		defer subscription.Unsubscribe()

		log.Printf("[Worker %d] started\n", id)

		select {}
	}

	for i := 0; i < WorkerCount; i++ {
		wg.Add(1)
		go worker(i + 1)
	}

	wg.Wait()
}
