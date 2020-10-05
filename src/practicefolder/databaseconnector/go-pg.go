package databaseconnector

import(
	"fmt"
	"os"
	pg "github.com/go-pg/pg"
)

func Connect() *pg.DB{
	opts := &pg.Options{
		User: "postgres",
		Password: "1234",
		Addr: "localhost:5432",
		Database: "practicefolder",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil{
		fmt.Println("Failed to connect to databse ")
		os.Exit(100)
	}
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