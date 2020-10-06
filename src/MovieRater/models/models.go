package models

import(
	"fmt"
	orm "github.com/go-pg/pg/orm"
	pg "github.com/go-pg/pg"
	db "MovieRater/databaseconnector"
)

type User struct{
	tablename struct{} `sql:"users"`
	Id int `sql:"id,pk,foreignkey" json:"id"`
	Name string `sql:"name" json:"name"`
	Phone string `sql:"phone" json:"phone"`
	Password  string `sql:"password" json:"password"`
	Email string `sql:"email,unique" json:"email"`

}

type Movie struct{
	tablename struct{} `sql:"movies"`
	Id int `sql:"id,pk,foreignkey" json:"id"`
	Name string `sql:"name" json:"name"`
	Genre string `sql:"genre" json:"genre"`
	Rating  string `sql:"rating" json:"rating"`
	Release_date string `sql:"release_date" json:"release_date"`
}

type Rating struct{
	tablename struct{} `sql:"movie_ratings"`
	Id int `sql:"id,pk" json:"id"`
	UserID int `sql:"user_id,pk" json:"user_id"`
	User *User `sql:"rel:has-one"`
	MovieID int `sql:"movie_id" json:"movie_id"`
	Movie *Movie `sql:"rel:has-one"`
	Rating  float64 `sql:"rating" json:"rating"`
}
func CreateProductItemsTable(db *pg.DB) error{
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&User{},opts)
	createErr2 := db.CreateTable(&Movie{},opts)
	createErr3 := db.CreateTable(&Rating{},opts)
	if createErr != nil{
		fmt.Println("Error while creating table user",createErr)
		return createErr
	}
	if createErr2 != nil{
		fmt.Println("Error while creating table movie",createErr2)
		return createErr2
	}
	if createErr2 != nil{
		fmt.Println("Error while creating table rating",createErr2)
		return createErr3
	}
	fmt.Println("Tables created successfully")
	return nil
}


func (u *User) CreateUser() error{
	dbpg := db.Connect()
	_,err:= dbpg.Model(u).Insert()
	return err
}

func GetAllUser() ([]User,error){
	dbpg := db.Connect()
	var model []User
	err := dbpg.Model(&model).Order("id ASC").Select()
	return model,err
}