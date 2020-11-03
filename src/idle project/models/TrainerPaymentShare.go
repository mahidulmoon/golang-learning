package models

import (
	// "fmt"
	"time"

	"github.com/UpskillBD/BE-TrainerDash/db"
)

type TrainerPaymentShare struct {
	tablename   struct{}  `pg:"trainerpaymentshare"`
	Id          int64     `pg:"id,pk" json:"id,omitempty"`
	Workshop_ID int64     `pg:"workshopid" json:"workshopid,omitempty"`
	Trainer_Id  int64     `pg:"trainer_id" binding:"required" json:"trainer_id,omitempty"`
	Share       float32   `pg:"share" binding:"required" json:"share"`
	Created     time.Time `pg:"created" json:"created,omitempty"`
	Updated     time.Time `pg:"updated" json:"updated,omitempty"`
}

func (s *TrainerPaymentShare) Add() error {
	s.Created = time.Now()
	_, err := db.DB.Model(s).Insert()
	return err
}

func GetAllShare() ([]TrainerPaymentShare, error) {
	var share []TrainerPaymentShare
	err := db.DB.Model(&share).Order("id ASC").Select()
	return share, err

}

func GetShareWid(id int) (TrainerPaymentShare, error) {
	var share TrainerPaymentShare
	err := db.DB.Model(&share).Where("workshop_id = ?", id).Select()
	return share, err
}

func (s *TrainerPaymentShare) UpdateShare() error {
	s.Updated = time.Now()
	_, err := db.DB.Model(s).Where("id = ?", s.Id).Update()
	return err
}

func DeleteShareid(id int) error {
	var share TrainerPaymentShare
	_, err := db.DB.Model(&share).Where("id = ?", id).Delete()
	return err
}
