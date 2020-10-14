package BE_TrainerDash

import "github.com/go-pg/pg/orm"
import "github.com/UpskillBD/BE-TrainerDash/models"
import db "github.com/UpskillBD/BE-TrainerDash/db"

func setupSchema() error {
	models := []interface{}{
		(*models.Files)(nil),
		(*models.WorkshopRequest)(nil),
		(*models.TrainerPayment)(nil),
	}

	for _, model := range models {
		err := db.DB.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

