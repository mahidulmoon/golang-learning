package models

import(
	
	"fmt"
	orm "github.com/go-pg/pg/orm"
	pg "github.com/go-pg/pg"
	db "practicefolder/databaseconnector"
	"time"
)

type ProductItem struct{
	RefPointer int `sql:"-"`
	tablename struct{} `sql:"product_items_collection"`
	ID int `sql:"id,pk" json:"id"`
	Name string `sql:"name,unique" json:"name"`
	Desc string `sql:"desc" json:"desc"`
	Image string `sql:"image" json:"image"`
	Price float64 `sql:"price,type:real" json:"price"`
	CreatedAt time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt time.Time `sql:"updated_at" json:"updated_at"`
	IsActive bool `sql:"is_active" json:"isactive"`
} 

func CreateProductItemsTable(db *pg.DB) error{
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&ProductItem{},opts)
	if createErr != nil{
		fmt.Println("Error while creating table products",createErr)
		return createErr
	}
	fmt.Println("Table productItems created successfully")
	return nil
}

func GetProducts()([]ProductItem,error){
	dbpg := db.Connect()
	var model []ProductItem
	err := dbpg.Model(&model).Order("id ASC").Select()
	return model,err
}

func (c *ProductItem) AddProduct() error{
	dbpg := db.Connect()
	//c.CreatedAt := time.now()
	_, err:= dbpg.Model(c).Insert()

	return err
}

func (c *ProductItem) UpdateProduct() error{
	dbpg := db.Connect()
	_,err := dbpg.Model(c).Where("id = ?",c.ID).Update()
	return err
}


func GetProductsById(id int)(ProductItem,error){
	dbpg := db.Connect()
	var model ProductItem
	err := dbpg.Model(&model).Where("id = ?",id).Select()
	return model,err
}

func DeleteProductsById(id int) error{
	dbpg := db.Connect()
	var model ProductItem
	models,err := dbpg.Model(&model).Where("id = ?",id).Delete()
	fmt.Println(models)
	return err
}