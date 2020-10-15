package BE_TrainerDash

import (
	"fmt"
	"github.com/UpskillBD/BE-TrainerDash/db"
	"github.com/UpskillBD/BE-TrainerDash/models"
	"github.com/go-pg/pg/orm"
	"testing"
)


func TestSetupSchema(t *testing.T) {

	models := []interface{}{
		(*models.Files)(nil),
		(*models.TrainerPayment)(nil),
		(*models.WorkshopRequest)(nil),
	}

	for _, model := range models {
		err := db.DB.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			fmt.Println("could not setup")
		}
	}

}