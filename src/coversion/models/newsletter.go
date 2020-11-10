package models

import (
	"coversion/db"
)

type NewsLetter struct {
	tableName struct{} `pg:"workshops"`
	ID        int      `pg:"id,pk" json:"id,omitempty"`
	Email     string   `pg:"email" json:"email"`
}

type NewsLetterService interface {
	SubscribeNewsLetter(email string) (*NewsLetter, error)
	GetNewsLetterEmails() ([]string, error)
}

func (n *NewsLetter) Add() error {
	_, err := db.DB.Model(n).Insert()
	return err
}

func (n *NewsLetter) ValidateNewLetter() (int, error) {
	var letter NewsLetter
	err := db.DB.Model(&letter).Where("email = ?", n.Email).Select()
	return letter.ID, err
}
