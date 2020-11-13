package db

import (
	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
)

var DB *pg.DB

func GetDB() (*pg.DB) {
	//dbString := "postgresql://upskill:zareen1234@webdb.cphjdxyzvku3.ap-southeast-1.rds.amazonaws.com:5432/upskill_webdb?sslmode=disable"
	dbString := "postgres://postgres:1234@localhost:5432/conversion?sslmode=disable" //localhost
	opt, err := pg.ParseURL(dbString)
	if err != nil {
		panic("could not set up database")
	}
	DB = pg.Connect(opt)
	return DB
}

