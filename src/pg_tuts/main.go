package main

import (
	"time"
	"fmt"
db	"pg_tuts/db"
pg "github.com/go-pg/pg"
)


func main(){
	fmt.Println("hello")
	pg_db := db.Connect()
	SaveProduct(pg_db)
}

func SaveProduct(dbRef *pg.DB){
	newPI := &db.ProductItem{
		Name: "Product 2",
		Desc: "Product 2 desc",
		Image: "this is image path new",
		Price: 4.5,
		Features: struct{
			Name string 
			Desc string 
			Imp int
		}{
			Name: "F2",
			Desc : "F2 desc",
			Imp: 3,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive: true,
	}
	//newPI.Save(dbRef)
	newPI.SaveAndReturn(dbRef)
}