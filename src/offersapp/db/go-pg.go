package db

import (
	"fmt"

	"github.com/go-pg/pg"
)

var DB *pg.DB

func init() {
	fmt.Println("gopg DB Initialized")

	opt, err := pg.ParseURL("postgres://postgres:1234@localhost:5432/tuts?sslmode=disable")
	if err != nil {
		panic(err)
	}

	DB = pg.Connect(opt)

}
