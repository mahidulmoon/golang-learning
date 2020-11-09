package coversion

import (
	"coversion/db"
	"coversion/models"
	"fmt"
	"testing"

	"github.com/go-pg/pg/orm"
)

func TestSetupSchema(t *testing.T) {

	models := []interface{}{
		(*models.User)(nil),
	}

	for _, model := range models {
		err := db.DB.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			fmt.Println("could not setup because: ", err)
		}
	}

}
