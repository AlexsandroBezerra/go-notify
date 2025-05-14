package dto

type CreateEmailRequest struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Priority  int16  `json:"priority"`
}

type CreateEmailResponse struct {
	ID string `json:"id"`
}
