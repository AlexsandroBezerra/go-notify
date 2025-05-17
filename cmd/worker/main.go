package main

import (
	"AlexsandroBezerra/go-notify/internal/queue/subject"
	"AlexsandroBezerra/go-notify/internal/queue/subscriber"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

const WorkerCount = 5

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	wg := sync.WaitGroup{}

	worker := func(id int) {
		defer wg.Done()
		handler := subscriber.NewEmailHandler(id)
		subscription, err := nc.QueueSubscribe(subject.Email, "email-queue", handler.ProcessMessage)
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
