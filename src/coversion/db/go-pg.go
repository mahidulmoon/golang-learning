package db

import (
	"fmt"

	"github.com/go-pg/pg"
)

var DB *pg.DB

func init() {
	fmt.Println("gopg DB Initialized")
	db_string := "postgres://postgres:1234@localhost:5432/conversion?sslmode=disable"
	opt, err := pg.ParseURL(db_string)
	if err != nil {
		panic(err)
	}

	DB = pg.Connect(opt)

}
