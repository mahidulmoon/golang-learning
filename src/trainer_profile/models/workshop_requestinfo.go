package models

import (
	"fmt"
	"github.com/UpskillBD/BE-TrainerDash/db"
	"time"
)

type FileResponse struct {
	File_id   int
	File_name string
	Link      string
}

type WorkshopRequest struct {
	tablename   struct{} `pg:"wsrequest_info"`
	Id          int64    `pg:"id,pk" json:"id,omitempty"`
	Workshop_ID int64    `pg:"workshopid,pk" json:"workshopid,omitempty"`
	Trainer_Id  int64    `pg:"trainer_id" binding:"required" json:"trainer_id"`
	Category    string   `pg:"category" binding:"required" json:"category"` //pg dosen't support this format as jsonb in pgdb
	//Category struct{} `pg:"category" binding:"required" json:"category"`    //for backup I commented here just remove the comment and run
	Description string `pg:"description,type:string" binding:"required" json:"description"`
	//Description struct{} `pg:"description,type:jsonb" binding:"required" json:"description"`
	Course_outline string                  `pg:"course_outline,type:jsonb" binding:"required" json:"course_outline"`
	Duration       int                     `pg:"duration" binding:"required" json:"duration"`
	Fees           float64                 `pg:"fees" json:"fees,omitempty" `
	Image          string                  `pg:"image" json:"image,omitempty"`
	Location       string                  `pg:"location" json:"location,omitempty"`
	Name           string                  `pg:"name" binding:"required" json:"name"`
	Status         string                  `pg:"status" binding:"required" default:"CREATED" json:"status,omitempty"`
	Workshop_ts    int                     `pg:"workshop_ts" binding:"required" json:"workshop_ts"`
	File           FileResponse `pg:"files,type:jsonb" json:"files"`

	Created time.Time `pg:"created" json:"created,omitempty"`
	Updated time.Time `pg:"updated" json:"updated,omitempty"`
}

func (w *WorkshopRequest) Add() error {

	w.Created = time.Now()
	_, err := db.DB.Model(w).Insert()
	return err
}

func GetWorkshopRequests() ([]WorkshopRequest, error) {

	var workshoprequest []WorkshopRequest
	err := db.DB.Model(&workshoprequest).Order("id ASC").Select()
	return workshoprequest, err

}

func GetWorkshopRequestById(id int) (WorkshopRequest, error) {
	var workshoprequest WorkshopRequest
	err := db.DB.Model(&workshoprequest).Where("id = ?", id).Select()
	return workshoprequest, err
}

func TrainerGetWorkshopRequestById(id int, tid int) (WorkshopRequest, error) {
	var workshoprequest WorkshopRequest
	err := db.DB.Model(&workshoprequest).Where("id = ?", id).Where("trainer_id=?",tid).Select()
	return workshoprequest, err
}

func DeleteWorkshopRequest(id int) error {
	var workshoprequest WorkshopRequest
	workshop, err := db.DB.Model(&workshoprequest).Where("id = ?", id).Delete()
	fmt.Println(workshop)
	return err
}

func (w *WorkshopRequest) Update() error {
	w.Updated = time.Now()
	_, err := db.DB.Model(w).Where("id = ?", w.Id).Update()
	return err
}

func GetWorkshopRequestbyStatus(status string)([]WorkshopRequest, error){
	var model []WorkshopRequest
	err:=db.DB.Model(&model).Where("status = ?", status).Select()
	return  model, err
}

func GetWSRbyWID(wid int)(WorkshopRequest,error){
var model WorkshopRequest
err:=db.DB.Model(&model).Where("workshopid = ?", wid).Select()
return  model, err
}

func TrainerGetWSRbyWID(wid int,tid int)(WorkshopRequest,error){
	var model WorkshopRequest
	err:=db.DB.Model(&model).Where("workshopid = ?", wid).Where("trainer_id=?",tid).Select()
	return  model, err
}

func GetWorkshopeRequestByTrainer(tid int64)([]WorkshopRequest,error){
	var model []WorkshopRequest
	err:=db.DB.Model(&model).Where("trainer_id = ?", tid ).Select()
	return  model, err
}

