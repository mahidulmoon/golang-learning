package models

import (
	"fmt"
	"testing"
	"time"
)

func TestGetWorkshops(t *testing.T) {
	ws, err := GetWorkshops()
	fmt.Println(len(ws), err)
	fmt.Println(ws[0].ID, ws[0].Name, ws[0].WorkshopTS, ws[0].Fee)
}

func TestGetWorkshopByID(t *testing.T) {
	ws, err := GetWorkshopByID(66)
	fmt.Println(ws, err)
}

func TestWorkshopAdd(t *testing.T) {
	ws := Workshop{
		tableName:     struct{}{},
		ID:            0,
		Name:          "Testing",
		Description:   "Description",
		CourseOutline: "Outline",
		Image:         "storage",
		YoutubeLink:   "youtube.com",
		Location:      04,
		Instructor:    8,
		Status:        "",
		WorkshopTS:    "",
		Duration:      "",
		Category:      "",
		Fee:           0,
		Type:          false,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
	err := ws.Add()
	if err != nil {
		fmt.Println("couldnot insert")
	}
}
