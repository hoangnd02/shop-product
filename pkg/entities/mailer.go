package entities

type MailerPayload struct {
	Key    string
	To     string
	Record map[string]interface{}
}
