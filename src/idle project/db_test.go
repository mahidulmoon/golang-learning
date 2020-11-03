package BE_TrainerDash

import (
	"fmt"
	"testing"

	"github.com/UpskillBD/BE-TrainerDash/db"
	"github.com/UpskillBD/BE-TrainerDash/models"
	"github.com/UpskillBD/BE-TrainerDash/api"
	"github.com/go-pg/pg/orm"
)

func TestSetupSchema(t *testing.T) {

	models := []interface{}{
		(*models.Files)(nil),
		(*models.TrainerPayment)(nil),
		(*models.WorkshopRequest)(nil),
		(*models.TrainerPaymentShare)(nil),
	}

	for _, model := range models {
		err := db.DB.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			fmt.Println("could not setup because: " ,err)
		}
	}

}

func TestToken(t *testing.T){
	token:="Bearer v2.local.2_Hxlwr29pnJsmdPhw42COtWFU-nGYOWFSWZJVy53J1B2-Ls1rQNEma2eTbgNJHeYHXnijalHPbQXDJcxeh92ASSWuROmc90sqJpUBlbkoVMK1MXtW8pV8YPnwIbtnMaFxEzVXpB16AF40ZTL_2b7mqhuwbYEzeT3mBZ5c375cCjFHJHpNe-cPGW1fIqnScAa5QLZg5KH9KLKuTOu92tp4vRs792UO3HPVql54rFL9ftE9j9835YvmvR2Jtivs5P5yps3V_zSmnOE9LpgweRDmJ7mP4UV31m6p9Uor42THZzVlxXRa1MNDsSxQ.ODc2OEUwQ0Q"
	userInfo, errToken := api.Instance.GetUserFromToken(token)

	fmt.Println(userInfo)
	fmt.Println(errToken)

}
