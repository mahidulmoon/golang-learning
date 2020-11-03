package models

import (
	"time"

	"github.com/UpskillBD/BE-TrainerDash/db"
)

type TrainerPayment struct {
	tablename  struct{}  `pg:"trainer_payment"`
	Id         int64     `pg:"id,pk" json:"id,omitempty"`
	Trainer_Id int64     `pg:"trainer_id" binding:"required" json:"trainer_id"`
	Bkash      string    `pg:"bkash" json:"bkash,omitempty"`
	BankInfo   string    `pg:"bank_info"json:"bank_info,omitempty"`
	Total      float32   `pg:"total" binding:"required" json:"total,omitempty" `
	Due        float32   `pg:"due" binding:"required" json:"due,omitempty"`
	Percentage float32   `pg:"percentage" binding:"required" json:"percentage"`
	Comments   string    `pg:"comments" json:"comments,omitempty"`
	Created    time.Time `pg:"created" json:"created,omitempty"`
	Updated    time.Time `pg:"updated" json:"updated,omitempty"`
}

func (c *TrainerPayment) Add() error {

	c.Created = time.Now()
	_, err := db.DB.Model(c).Insert()
	return err
}

func GetAllTrainerPayments() ([]TrainerPayment, error) {

	var models []TrainerPayment
	err := db.DB.Model(&models).Order("id ASC").Select()
	return models, err

}

func GetTrainerPaymentById(id int) (TrainerPayment, error) {
	var model TrainerPayment
	err := db.DB.Model(&model).Where("id = ?", id).Select()
	return model, err
}

func GetTrainersPaymentDue() ([]TrainerPayment, error) {
	var models []TrainerPayment
	err := db.DB.Model(&models).Where("due > 0").Select()
	return models, err
}

func DeleteTrainerPayment(id int) error {
	var model TrainerPayment
	_, err := db.DB.Model(&model).Where("id = ?", id).Delete()

	return err
}

func (c *TrainerPayment) Update() error {
	c.Updated = time.Now()
	_, err := db.DB.Model(c).Where("id = ?", c.Id).Update()
	return err
}

func (c *TrainerPayment) TrainerUpdate() error {
	c.Updated = time.Now()
	c.Comments = "User has updated payment info" + c.Comments
	_, err := db.DB.Model(c).Where("id = ?", c.Id).Where("trainer_id=?", c.Trainer_Id).Update()
	return err
}
func GetTrainerPaymentByTID(tid int) (TrainerPayment, error) {
	var model TrainerPayment
	err := db.DB.Model(&model).Where("trainer__id = ?", tid).Limit(1).Select()
	return model, err
}

//type PaymentDash struct{
//	WorkshopId int64
//	WorkshopName string
//	WorkshopDate string
//	Registrants []OldBackend.Registrants
//	Percentage float32
//	RawTotal float64
//	CalculatedTotal float64
//
//}

func TrainerpaymentDash(courseid int) {

}
