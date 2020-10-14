package models

import (
	"time"
	"github.com/UpskillBD/BE-TrainerDash/db"
)

type Files struct {
	tablename    struct{}  `pg:"files"`
	Id           int64     `pg:"id,pk" json:"id,omitempty"`
	Trainer_Id   int64     `pg:"trainer_id" binding:"required" json:"trainer_id"`
	Role         string    `pg:"role" binding:"required" json:"role"`
	OriginalName string    `pg:"originalname" binding:"required" json:"originalname"`
	Type         string    `pg:"type" binding:"required" json:"type" `
	Name         string    `pg:"name" binding:"required" json:"name"`
	Link         string    `pg:"link" json:"link,omitempty"`
	Access       string    `pg:"access" binding:"required" json:"access"`
	Created      time.Time `pg:"created" json:"created,omitempty"`
	Updated      time.Time `pg:"updated" json:"updated,omitempty"`
}

func (f *Files) Add() error {
	f.Created = time.Now()
	_, err := db.DB.Model(f).Insert()
	return err
}

func DeleteFile(id int) error {
	var file Files
	_, err := db.DB.Model(&file).Where("id = ?", id).Delete()
	return err
}

func GetAllFiles() ([]Files, error) {
	var file []Files
	err := db.DB.Model(&file).Select()
	return file, err
}

func GetFilesByID(id int) (Files, error) {

	var file Files
	err := db.DB.Model(&file).Where("id = ?", id).Select()
	return file, err
}

func (f *Files) Update() error {

	f.Updated = time.Now()
	_, err := db.DB.Model(f).Where("id = ?", f.Id).Update()
	return err
}
