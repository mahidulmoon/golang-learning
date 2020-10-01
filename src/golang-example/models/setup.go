package models
import (
	// "fmt"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"
	//"database/sql"
  	"fmt"

  	_ "github.com/lib/pq"

   )
func SetupModels() *gorm.DB {//db, err := gorm.Open("sqlite3", "test.db")// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()// To get the value from the config file using key// viper package read .env
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1234"
		dbname   = "medium"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)

	  
	//fmt.Println("conname is\t\t", psqlInfo)
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
	 panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Book{})// Initialise value
	m := Book{Author: "author1", Title: "title1"}
	db.Create(&m)
	return db
   }