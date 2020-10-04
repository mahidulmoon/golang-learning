package db

import (
	"fmt"
	"github.com/go-pg/pg"
)

type Params struct{
	Param1 string
	Param2 string
}

func PlaceHolderDemo(db *pg.DB) error{
	var value string
	params := Params{
		Param1: "this is param 1",
		Param2: "this is param 2",
	}
	var query string = "Select ?"
	_, selectErr:=db.Query(pg.Scan(&value),query,params)
	if selectErr != nil{
		fmt.Println("error while running the select query")
		return selectErr
	}
	fmt.Println("Scan successful")
	return nil
}