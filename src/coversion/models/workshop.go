package models

import (
	"coversion/db"
	"fmt"
	"time"
)

//WorkShop holds the workshop information
type Workshop struct {
	tableName     struct{}  `pg:"workshops"`
	ID            int       `pg:"id,pk" json:"id,omitempty"`
	Name          string    `pg:"name" json:"name"`
	Description   string    `pg:"description" json:"description"`
	CourseOutline string    `pg:"course_outline" json:"course_outline"`
	Image         string    `pg:"image" json:"image"`
	YoutubeLink   string    `pg:"youtube_link" json:"youtube_link"`
	Location      string    `pg:"location" json:"location"`
	Instructor    int       `pg:"instructor" json:"instructor"`
	Status        string    `pg:"status"  json:"status"`
	WorkshopTS    string    `pg:"workshop_ts" json:"workshop_ts"`
	Duration      string    `pg:"duration" json:"duration"`
	Category      string    `pg:"category" json:"category"`
	Fee           int       `pg:"fee" json:"fee"`
	Type          bool      `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (j *Workshop) Add() error {
	j.CreatedAt = time.Now()
	_, err := db.DB.Model(j).Insert()
	return err
}

func (j *Workshop) Update() error {
	res, err := db.DB.Model(j).Where("id = ?", j.ID).Update()
	fmt.Println(res)
	return err
}

func GetWorkshops() ([]Workshop, error) {
	var models []Workshop
	err := db.DB.Model(&models).Select()
	return models, err
}

func GetWorkshopByID(id int) ([]Workshop, error) {
	var jobs []Workshop
	err := db.DB.Model(&jobs).Order("id ASC").Where("id = ?", id).Select()
	return jobs, err
}
