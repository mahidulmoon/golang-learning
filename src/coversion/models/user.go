package models

import (
	"coversion/db"
	"time"
)

type User struct {
	tablename       struct{}  `pg:"user"`
	ID              int64     `json:"id" pg:"id,pk"`
	Name            string    `pg:"name" `
	Email           string    `pg:"email,unique" binding:"required"`
	Password        string    `json:"password"  pg:"password"`
	Phone           string    `pg:"phone,unique" binding:"required" json:"phone"`
	EmailVerified   bool      `json:"email_verfied" pg:"email_verified"`
	Token           string    `json:"token,omitempty"`
	EmailVerifiedAt string    `json:"email_verified_at,omitempty" pg:"email_verified_at"`
	CreatedAt       time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" pg:"updated_at"`
}

func (u *User) Add() error {
	u.CreatedAt = time.Now()
	_, err := db.DB.Model(u).Insert()

	return err
}

func GetUserById(id int) (User, error) {
	var user User
	err := db.DB.Model(&user).Where("id = ?", id).Select()
	return user, err
}

func (u *User) Validate() (int64, error) {
	var user User
	err := db.DB.Model(&user).Where("email = ?", u.Email).WhereOr("phone = ?", u.Phone).Select()
	return user.ID, err
}
