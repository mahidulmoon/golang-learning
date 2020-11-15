package models

import (
	"fmt"
	"time"

	db "github.com/UpskillBD/upskill-main/db"
)

const (
	DRAFT    = 1
	PRIVATE  = 2
	UNLISTED = 3
	LIVE     = 4
	UPCOMING = 5
)

//WorkShop holds the workshop information
type Workshop struct {
	tableName     struct{}  `json:"-" sql:"workshops"`
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	CourseOutline string    `json:"course_outline"`
	Image         string    `json:"image"`
	YoutubeLink   string    `json:"youtube_link"`
	Location      int       `json:"location"`
	Instructor    int       `json:"instructor"`
	Status        string    `json:"status"`
	WorkshopTS    string    `json:"workshop_ts"`
	Duration      string    `json:"duration"`
	Category      string    `json:"category"`
	Fee           int       `json:"fee"`
	Type          bool      `json:"type"`
	CreatedBy     int       `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (j *Workshop) Add() error {
	j.CreatedAt = time.Now()
	res, err := db.GetDB().Model(j).Insert()
	fmt.Println(res)
	return err
}

func (j *Workshop) Update() error {
	res, err := db.GetDB().Model(j).Where("id = ?", j.ID).Update()
	fmt.Println(res)
	return err
}

func GetPaginatedWorkshops(num int) ([]Workshop, error) {
	var models []Workshop
	offset := 0
	if num > 1 {
		offset += 5 * (num - 1)
	}
	err := db.GetDB().Model(&models).Limit(4).Offset(offset).Select()
	return models, err
}

func GetWorkshops() ([]Workshop, error) {
	var models []Workshop
	err := db.GetDB().Model(&models).Select()
	return models, err
}

func GetWorkshopByID(id int) ([]Workshop, error) {
	var models []Workshop
	err := db.GetDB().Model(&models).
		Order("id ASC").
		Where("id = ?", id).
		Select()
	return models, err
}

func WSJoinReq(id int) error {
	var j *Workshop
	joined := "joined"
	updatedAt := time.Now()
	_, err := db.GetDB().Model(j).Set("status = ?", joined).Set("updated_at = ?", updatedAt).Where("id = ?", id).Update()
	return err
}

func WSCancelReq(id int) error {
	var j *Workshop
	cancel := "cancel"
	updatedAt := time.Now()
	_, err := db.GetDB().Model(j).Set("status = ?", cancel).Set("updated_at = ?", updatedAt).Where("id = ?", id).Update()
	return err
}

func GetWorkshopStatusByID(id int) (string, error) {
	var status string
	err := db.GetDB().Model((*Workshop)(nil)).Column("status").Where("id = ?", id).Select(&status)
	return status, err
}

// func GetWorkshopStatusByID(id int) (string,error) {
// 	var model Workshop
// 	err := db.GetDB().Model(&model).Where("id = ?", id).Select()
// 	return model.Status,err
// }
