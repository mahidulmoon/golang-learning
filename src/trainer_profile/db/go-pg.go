package db

import (
	"fmt"
	cfg "github.com/UpskillBD/BE-TrainerDash/config"
	"github.com/go-pg/pg"
)

var DB *pg.DB

func init() {
	fmt.Println("gopg DB Initialized")
	cfg.ViperSetup()
	db_string := cfg.Config.DbString
	opt, err := pg.ParseURL(db_string)
	if err != nil {
		panic(err)
	}

	DB = pg.Connect(opt)

}

