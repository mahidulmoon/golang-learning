package db
import(
	"fmt"
	"time"
	orm "github.com/go-pg/pg/orm"
	pg "github.com/go-pg/pg"
)


type ProductItem struct{
	RefPointer int `sql:"-"`
	tablename struct{} `sql:"product_items_collection"`
	ID int `sql:"id,pk"`
	Name string `sql:"name,unique"`
	Desc string `sql:"desc"`
	Image string `sql:"image"`
	Price float64 `sql:"price,type:real"`
	Features struct{
		Name string
		Desc string
		Imp int
	} `sql:"features,type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive bool `sql:"is_active"`
}

func (pi *ProductItem) Save(db *pg.DB) error{
	insertErr := db.Insert(pi)
	if insertErr != nil{
		fmt.Println("Error while inserting new item into DB")
		return insertErr
	}
	fmt.Println("ProductItem inserted successfull")
	return nil
}

func CreateProductItemsTable(db *pg.DB) error{
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&ProductItem{},opts)
	if createErr != nil{
		fmt.Println("Error while creating table productItems")
		return createErr
	}
	fmt.Println("Table productItems created successfully")
	return nil
}