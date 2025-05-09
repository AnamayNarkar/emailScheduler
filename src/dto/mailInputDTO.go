package dto

type MailInputDTO struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	SendAt  string `json:"sendAt"`
}
