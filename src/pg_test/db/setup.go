package db


import (
    "fmt"

    "github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
)


func SetupModels() *pg.DB{
    fmt.Println("success")
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "1234",
		Database: "pgtest",
	})
	err := createSchema(db)
    if err != nil {
        panic(err)
    }
	user1 := &User{
        Name:   "admin",
        Email: "admin2@admin",
    }
    _, err = db.Model(user1).Insert()
    if err != nil {
        panic(err)
    }
	return db
}
func createSchema(db *pg.DB) error {
    fmt.Println("inside schema")
    models := []interface{}{
        (*User)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            //Temp: true,
        })
        if err != nil {
            return err
        }
    }
    return nil
}