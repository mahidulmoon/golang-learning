package models

import (
	"time"
)

type User struct {
	tablename       struct{}  `pg:"user"`
	ID              int64     `json:"id" pg:"id,pk"`
	Name            string    `pg:"name" binding:"required"`
	Email           string    `pg:"email" binding:"required"`
	Password        string    `json:"password" binding:"required" pg:"password"`
	Phone           string    `pg:"phone" binding:"required" json:"phone"`
	EmailVerified   bool      `json:"email_verfied" pg:"email_verified"`
	Token           string    `json:"token,omitempty"`
	EmailVerifiedAt string    `json:"email_verified_at,omitempty" pg:"email_verified_at"`
	CreatedAt       time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" pg:"updated_at"`
}
