package models

import(
	//"fmt"
	"time"
	//orm "github.com/go-pg/pg/orm"
	//pg "github.com/go-pg/pg"
	db "github.com/UpskillBD/trainer-dashboard-backend/db"
)



type Files struct{
	tablename struct{} `pg:"files"`
	Id int64 `pg:"id,pk" json:"id,omitempty"` 
	Trainer_Id int64 `pg:"trainer_id" binding:"required" json:"trainer_id"`
	Role string `pg:"role" binding:"required" json:"role"`
	OriginalName string `pg:"originalname" binding:"required" json:"originalname"`
	Type string `pg:"type" binding:"required" json:"type" `
	Name string `pg:"name" binding:"required" json:"name"`
	Link string `pg:"link" json:"link,omitempty"`
	Access string `pg:"access" binding:"required" json:"access"`
	Created time.Time `pg:"created" json:"created,omitempty"`
	Updated time.Time `pg:"updated" json:"updated,omitempty"`
}

// func CreateFilesTable(db *pg.DB) error {
// 	opts := &orm.CreateTableOptions{
// 		IfNotExists: true,
// 	}
// 	createErr := db.CreateTable(&Files{}, opts)
// 	if createErr != nil {
// 		fmt.Println("Error while creating table Files", createErr)
// 		return createErr
// 	}
// 	fmt.Println("Tables created successfully")
// 	return nil
// }


func(f *Files) AddFiles() error {
	dbpg := db.Connect()
	f.Created=time.Now()
	_, err := dbpg.Model(f).Insert()
	return err
}