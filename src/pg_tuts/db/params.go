package db

import (
	"fmt"
	"github.com/go-pg/pg"
)

func PlaceHolderDemo(db *pg.DB) error{
	var value int
	var query string = "Select ?"
	_, selectErr:=db.Query(pg.Scan(&value),query,42)
	if selectErr != nil{
		fmt.Println("error while running the select query")
		return selectErr
	}
	fmt.Println("Scan successful")
	return nil
}