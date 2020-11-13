package UpskillBD

import (
	"github.com/UpskillBD/upskill-main/db"
	"github.com/UpskillBD/upskill-main/models"
	"fmt"
	"testing"

	"github.com/go-pg/pg/orm"
)

func TestSetupSchema(t *testing.T) {

	models := []interface{}{
		(*models.User)(nil),
		(*models.Workshop)(nil),
		(*models.NewsLetter)(nil),
	}

	for _, model := range models {
		err := db.GetDB().Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			fmt.Println("could not setup because: ", err)
		}
	}

}