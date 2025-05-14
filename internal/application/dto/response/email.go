package response

import "time"

type CreateEmail struct {
	ID string `json:"id"`
}

type Email struct {
	ID        string    `json:"id"`
	Recipient string    `json:"recipient"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	Priority  int16     `json:"priority"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type ListEmail []Email
