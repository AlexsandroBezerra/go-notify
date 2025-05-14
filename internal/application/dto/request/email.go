package request

type CreateEmail struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Priority  int16  `json:"priority"`
}
