package models

type NewsLetter struct {
	ID    int    `db:"id"`
	Email string `json:"email"`
}

type NewsLetterService interface {
	SubscribeNewsLetter(email string) (*NewsLetter, error)
	GetNewsLetterEmails() ([]string, error)
}
