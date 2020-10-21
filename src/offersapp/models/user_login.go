package models

import (
	"fmt"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

type Login struct {
	tablename     struct{} `pg:"login"`
	Id            int64    `pg:"id,pk" json:"id,omitempty"`
	Email         string   `pg:"email,unique" json:"email"`
	Password_hash string   `pg:"password_hash" binding:"required" json:"password_hash"`
	Password      string   `pg:"password" json:"password"`
}

func CreateLoginTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Login{}, opts)
	if createErr != nil {
		fmt.Println("Error while creating login products", createErr)
		return createErr
	}
	fmt.Println("Table login created successfully")
	return nil
}
