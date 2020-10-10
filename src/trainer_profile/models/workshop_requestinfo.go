package models

import (
	"fmt"
	"time"

	db "github.com/UpskillBD/trainer-dashboard-backend/db"
	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

type TypeWorkshopRequestFile struct {
	File_id   int
	File_name string
	Link      string
}

type WorkshopRequest struct {
	tablename   struct{} `pg:"workshop_request_info"`
	Id          int64    `pg:"id,pk" json:"id,omitempty"`
	Workshop_ID int64    `pg:"workshopid,pk" json:"workshopid,omitempty"`
	Trainer_Id  int64    `pg:"trainer_id" binding:"required" json:"trainer_id"`
	Category    string   `pg:"category" binding:"required" json:"category"` //pg dosen't support this format as jsonb in pgdb
	//Category struct{} `pg:"category" binding:"required" json:"category"`    //for backup I commented here just remove the comment and run
	Description string `pg:"description,type:string" binding:"required" json:"description"`
	//Description struct{} `pg:"description,type:jsonb" binding:"required" json:"description"`
	Course_outline string                  `pg:"course_outline,type:jsonb" binding:"required" json:"course_outline"`
	Duration       int                     `pg:"duration" binding:"required" json:"duration"`
	Fees           float64                 `pg:"fees" binding:"required" json:"fees" `
	Image          string                  `pg:"image" binding:"required" json:"image"`
	Location       string                  `pg:"location" json:"location"`
	Name           string                  `pg:"name" binding:"required" json:"name"`
	Status         string                  `pg:"status" binding:"required" default:"CREATED" json:"status"`
	Workshop_ts    int                     `pg:"workshop_ts" binding:"required" json:"workshop_ts"`
	File           TypeWorkshopRequestFile `pg:"files,type:jsonb" json:"files"`

	// File struct{
	// 	File_id int
	// 	File_name string
	// 	Link string
	// } `pg:"file,type:jsonb"`

	Created time.Time `pg:"created" json:"created,omitempty"`
	Updated time.Time `pg:"updated" json:"updated,omitempty"`
}

func CreateWorkShopRequestTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&WorkshopRequest{}, opts)
	createErr2 := db.CreateTable(&Files{}, opts)
	if createErr != nil {
		fmt.Println("Error while creating table WorkshopRequest", createErr)
		return createErr
	}
	if createErr2 != nil {
		fmt.Println("Error while creating table Files", createErr2)
		return createErr2
	}
	fmt.Println("Tables created successfully")
	return nil
}

func (w *WorkshopRequest) AddWorkShop() error {
	dbpg := db.Connect()
	w.Created = time.Now()
	_, err := dbpg.Model(w).Insert()
	return err
}

func GetWorkshopRequest() ([]WorkshopRequest, error) {
	dbpg := db.Connect()
	var workshoprequest []WorkshopRequest
	err := dbpg.Model(&workshoprequest).Order("id ASC").Select()
	return workshoprequest, err

}

func GetWorkshopRequestById(id int) (WorkshopRequest,error){
	dbpg := db.Connect()
	var workshoprequest WorkshopRequest
	err := dbpg.Model(&workshoprequest).Where("id = ?",id).Select()
	return workshoprequest,err
}

func DeleteWorkshopRequestById(id int) error{
	dbpg := db.Connect()
	var workshoprequest WorkshopRequest
	workshop,err := dbpg.Model(&workshoprequest).Where("id = ?",id).Delete()
	fmt.Println(workshop)
	return err
}

func (w *WorkshopRequest) UpdateWorkShop() error {
	dbpg := db.Connect()
	w.Updated = time.Now()
	_, err := dbpg.Model(w).Where("id = ?",w.Id).Update()
	return err
}