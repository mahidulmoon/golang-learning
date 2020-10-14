package db

import (
	cfg "github.com/UpskillBD/BE-TrainerDash/config"
	"fmt"
	"github.com/go-pg/pg"
)

var DB *pg.DB

func init() {
	fmt.Println("gopg DB Initialized")
	fmt.Println(cfg.Config.DbName)
	fmt.Println(cfg.Config.DbString)
	db_string := cfg.Config.DbName
	opt, err := pg.ParseURL(db_string)
	if err != nil {
		fmt.Println(err)
	}
	DB = pg.Connect(opt)
}

//func isDBAlive() error {
//	ctx := context.Background()
//	err := DB.Ping(ctx)
//	return err
//}