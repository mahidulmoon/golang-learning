package db

import (
	"fmt"
	"os"

	pg "github.com/go-pg/pg"
)

var db *pg.DB

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "1234",
		Addr:     "localhost:5432",
		Database: "testdb",
	}
	db = pg.Connect(opts)
	if db == nil {
		fmt.Println("Failed to connect to databse ")
		os.Exit(100)
	}
	// opts, err := pg.ParseURL("postgresql://mahid:m@hid4567@webdb.cphjdxyzvku3.ap-southeast-1.rds.amazonaws.com:5432/testdb?sslmode=disable")
	// db := pg.Connect(opts)
	// if err != nil {
	// 	fmt.Println("Failed to connect to databse ")
	// 	os.Exit(100)
	// }
	fmt.Println("Connection to database successful")
	//CreateProductItemsTable(db)
	// closeErr := db.Close()
	// if closeErr != nil{
	// 	fmt.Println("Error while closing the connection")
	// 	os.Exit(100)
	// }
	// fmt.Println("Connection close successfully")
	return db
}
