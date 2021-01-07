package models

import (
	"encoding/json"
	"github.com/Upskillbd/newcsv/db"
	"database/sql"
	"fmt"
	"reflect"
)

type QueryStruct struct {
	Query string `json:"query" binding:"required"`
}

func GetRows(query_string string) []byte {
	if db.DB != nil {
		// Query contains no user-defined variables [Safe]
		query := fmt.Sprintf(query_string)
		//rows, err := db.DB.Query(query)
		rows, err := queryToJson(db.DB, query)
		if err != nil {
			panic(err)
		}
		return rows
	}
	return nil
}

func queryToJson(db *sql.DB, query string, args ...interface{}) ([]byte, error) {
	// an array of JSON objects
	// the map key is the field name
	var objects []map[string]interface{}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		// figure out what columns were returned
		// the column names will be the JSON object field keys
		columns, err := rows.ColumnTypes()
		if err != nil {
			return nil, err
		}

		// Scan needs an array of pointers to the values it is setting
		// This creates the object and sets the values correctly
		values := make([]interface{}, len(columns))
		object := map[string]interface{}{}
		for i, column := range columns {
			v := reflect.New(column.ScanType()).Interface()
			switch v.(type) {
			case *[]uint8:
				v = new(string)
			default:
				// use this to find the type for the field
				// you need to change
				// log.Printf("%v: %T", column.Name(), v)
			}

			object[column.Name()] = v
			values[i] = object[column.Name()]
		}

		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		objects = append(objects, object)
	}

	// indent because I want to read the output
	return json.MarshalIndent(objects, "", "\t")
}