package models

import (
	"fmt"
	"testing"
)

func TestAddWorkshopRequest(t *testing.T) {
	fmt.Println("test is working")
	w := WorkshopRequest{

		Trainer_Id:     40,
		Workshop_ID:    100,
		Category:       "nothing ",
		Description:    "big description",
		Course_outline: "i don't know ",
		Duration:       60,
		Fees:           1400.00,
		Image:          "nil",
		Location:       "double check location",
		Name:           "mahidulmoon",
		Status:         "modified",
		Workshop_ts:    40,
		File: struct {
			File_id   int
			File_name string
			Link      string
		}{
			File_id:   2,
			File_name: "file name",
			Link:      "facebook.com",
		},
	}
	err := w.Add()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	fmt.Println("test passed workshop added.")

}

func TestGetWorkshopRequest(t *testing.T) {
	w, err := GetWorkshopRequests()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	fmt.Println(w, err)

}

func TestGetWSbyWID(t *testing.T) {
	w, err := TrainerGetWSRbyWID(100,40)
	//w,err:=GetWorkshopRequestByTrainer(40)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	fmt.Println(w, err)

}
