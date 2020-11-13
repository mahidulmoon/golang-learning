package models

import (
	db "github.com/UpskillBD/upskill-main/db"
)

type NewsLetter struct {
	ID    int     `db:"id"`
	Email string `json:"email"`
}

type NewsLetterService interface {
	SubscribeNewsLetter(email string) (*NewsLetter, error)
	GetNewsLetterEmails() ([]string, error)
}

func (n *NewsLetter) Add() error {
	_, err := db.GetDB().Model(n).Insert()
	return err
}

func (n *NewsLetter) ValidateNewLetter() (int, error) {
	var letter NewsLetter
	err := db.GetDB().Model(&letter).Where("email = ?", n.Email).Select()
	return letter.ID, err
}