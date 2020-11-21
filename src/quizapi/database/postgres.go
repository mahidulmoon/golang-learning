package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import solely for its initialization
)

const (
	DriverName = "postgres"
)

var db *sql.DB

func init() {
	//---Loads all the variables in the .env file into the OS environment
	err := godotenv.Load()
	fmt.Println("db initialized")
	//---Gets the value for the DB_CRED environment variable
	DataSource := os.Getenv("DB_CRED")
	fmt.Println("data source is " + DataSource)
	database, err := sql.Open(DriverName, DataSource)
	if err != nil {
		fmt.Println("db init error")
		panic(err)

	}
	db = database
}

func Get() *sql.DB {
	return db
}

func Close() {
	db.Close()
}

func GetRows(model string) *sql.Rows {
	if db != nil {
		// Query contains no user-defined variables [Safe]
		query := fmt.Sprintf("SELECT * FROM %s", model)
		rows, err := db.Query(query)
		if err != nil {
			panic(err)
		}
		return rows
	}
	return nil
}

func GetRowByDCode(model string, code int) *sql.Rows {
	if db != nil {
		// Query contains no user-defined variables [Safe]
		query := fmt.Sprintf("SELECT * FROM %s WHERE d_id = $1", model)
		rows, err := db.Query(query, code)
		if err != nil {
			panic(err)
		}
		return rows
	}
	return nil
}

func DeleteRowById(model string, id int) (sql.Result, error) {
	if db != nil {
		// Query contains no user-defined variables [Safe]
		query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", model)

		res, err := db.Exec(query, id)
		if err != nil {
			return nil, err
		}

		count, err := res.RowsAffected()
		if err != nil {
			return nil, err
		}

		if count == 0 {
			return nil, errors.New("No data found!")
		}
		//fmt.Println(count)

		return res, nil
	}
	return nil, errors.New("Database needs to be initialized!")
}
func GetRowById(model string, id int) *sql.Row {
	if db != nil {
		// Query contains no user-defined variables [Safe]
		query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", model)
		row := db.QueryRow(query, id)
		return row
	}
	return nil
}
