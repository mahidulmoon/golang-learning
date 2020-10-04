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
	//SaveProduct(pg_db)
	//db.PlaceHolderDemo(pg_db)
	//DeleteItem(pg_db)
	// UpdateItemPrice(pg_db)
	GetByID(pg_db)

}

func SaveProduct(dbRef *pg.DB){
	newPI := &db.ProductItem{
		Name: "Product 8",
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
	// newPI2 := &db.ProductItem{
	// 	Name: "Product 7",
	// 	Desc: "Product 2 desc",
	// 	Image: "this is image path new",
	// 	Price: 4.5,
	// 	Features: struct{
	// 		Name string 
	// 		Desc string 
	// 		Imp int
	// 	}{
	// 		Name: "F2",
	// 		Desc : "F2 desc",
	// 		Imp: 3,
	// 	},
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// 	IsActive: true,
	// }
	newPI.Save(dbRef)
	//newPI.SaveAndReturn(dbRef)
	// totalItems := []*db.ProductItem{newPI1,newPI2}
	// newPI1.SaveMultiple(dbref,totalItems)
}

func DeleteItem(dbRef *pg.DB){
	newPI := &db.ProductItem{
		Name: "Product 8",
	}
	newPI.DeleteItem(dbRef)
}

func UpdateItemPrice(dbref *pg.DB){
	newPI := &db.ProductItem{
		ID: 2,
		Price: 2.0,
	}
	newPI.UpdatePrice(dbref)
}

func GetByID(dbRef *pg.DB){
	newPI := &db.ProductItem{
		ID: 1,
		Name: "Product 1",
	}
	newPI.GetByID(dbRef)
}