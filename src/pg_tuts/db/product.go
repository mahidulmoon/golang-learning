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
	ID int `sql:"id,pk" json:"id,omitempty"`
	Name string `sql:"name,unique"`
	Desc string `sql:"desc"`
	Image string `sql:"image" json:"image"`
	Price float64 `sql:"price,type:real" json:"price"`
	Features struct{
		Name string
		Desc string
		Imp int
	} `sql:"features,type:jsonb"`
	CreatedAt time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt time.Time `sql:"updated_at" json:"updated_at"`
	IsActive bool `sql:"is_active" json:"is_active"`
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
func (pi *ProductItem) DeleteItem(db *pg.DB) error{
	_, deleteErr:= db.Model(pi).Where("name = ?name").Delete()
	if deleteErr !=nil{
		fmt.Println("Error while deleting item")
		return deleteErr
	}
	fmt.Println("Delete successful")
	return nil
}
func (pi *ProductItem) GetByID(db *pg.DB) error{
	getErr := db.Model(pi).Where("id = ?0",pi.ID).Where("name = ?0",pi.Name).Select()
	if getErr != nil{
		fmt.Println("Error while getting value but id.")
		return getErr
	}
	fmt.Println("Get by ID successfull")
	return nil
}

func (pi *ProductItem) SaveAndReturn(db *pg.DB) (*ProductItem,error){
	InsertResult,insertErr := db.Model(pi).Returning("+").Insert()
	if insertErr !=nil{
		fmt.Println("Error while inserting new item",insertErr)
		return nil,insertErr
	}
	fmt.Println("ProductItem Inserted successfull",InsertResult)
	return pi,nil

}
func (pi *ProductItem) SaveMultiple(db *pg.DB,items ...*ProductItem) error{
	_, insertErr:= db.Model(items[0],items[1]).Insert()
	if insertErr != nil{
		fmt.Println("error while inserting bulk items")
		return insertErr
	}
	fmt.Println("Bulk insert successfully")
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

func (pi *ProductItem) UpdatePrice(db *pg.DB) error{
	_, updateErr:= db.Model(pi).Set("price = ?price").Where("id = ?id").Update()
	if updateErr != nil{
		fmt.Println("Error while updating price")
		return updateErr
	}
	fmt.Println("Price updated successfully")
	return nil
}

