package main

import(
	//"fmt"
	db "MovieRater/databaseconnector"
	"MovieRater/models"
)

func main(){
	//fmt.Println("hello")
	dbpg := db.Connect()
	models.CreateProductItemsTable(dbpg)
}