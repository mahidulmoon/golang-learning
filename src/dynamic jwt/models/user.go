package models

import (
	"time"

	db "github.com/UpskillBD/upskill-main/db"
)

//User holds the user information
type User struct {
	tablename       struct{}  `pg:"user"`
	ID              int64     `json:"id" pg:"id,pk"`
	Name            string    `pg:"name" json:"name"`
	Email           string    `pg:"email,unique" json:"email" binding:"required"`
	Password        string    `json:"password"  pg:"password"`
	Phone           string    `pg:"phone,unique" binding:"required" json:"phone"`
	EmailVerified   bool      `json:"email_verfied" pg:"email_verified"`
	Token           string    `json:"token,omitempty"`
	EmailVerifiedAt string    `json:"email_verified_at,omitempty" pg:"email_verified_at"`
	CreatedAt       time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" pg:"updated_at"`
}

// type UserService interface {
// 	GetUserByEmail(email string) (*User, error)
// 	GetUserByID(userID string) (*User, error)
// 	GetUserByEmailAndPhone(email, phone string) (*User, error)
// 	GetUsers() ([]User, error)
// 	CreateUser(user *User) (*User, error)
// 	UpdateUser(user *User) (*User, error)
// }

func (u *User) Add() error {
	u.CreatedAt = time.Now()
	_, err := db.GetDB().Model(u).Insert()

	return err
}

func (u *User) Validate() (int64, error) {
	var user User
	err := db.GetDB().Model(&user).Where("email = ?", u.Email).WhereOr("phone = ?", u.Phone).Select()
	return user.ID, err
}

func GetUserById(id int) (User, error) {
	var user User
	err := db.GetDB().Model(&user).Where("id = ?", id).Select()
	return user, err
}

func (u *User) GetUserfromToken(token string) (int64, error) {
	var user User
	err := db.GetDB().Model(&user).Where("token = ?", token).Select()
	return user.ID, err
}

func (u *User) Authenticate(email string, password string) bool {
	var user User
	err := db.GetDB().Model(&user).Where("email = ?", email).Where("password = ?", password).Select()
	if err != nil {
		return false
	} else {
		return true
	}
}
