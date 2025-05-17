package message

type Email struct {
	ID        string `json:"id"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Priority  int16  `json:"priority"`
}
