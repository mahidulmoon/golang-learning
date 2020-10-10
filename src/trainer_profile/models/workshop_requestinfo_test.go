package models

import (
	"fmt"
	"testing"
)

func TestAddWorkshopRequest(t *testing.T) {
	fmt.Println("test is working")
	w := WorkshopRequest{
		
		Trainer_Id : 4,
		Workshop_ID : 10,
		Category : "nothing learning",
		Description :" description",
		Course_outline : "i do know ",
		Duration : 60,
		Fees : 1400.00,
		Image : "nil",
		Location :"check location",
		Name :"mahidul",
		Status : "modified",
		Workshop_ts : 40,
		File : struct{
			File_id   int
			File_name string
			Link      string
		}{
			File_id : 2,
			File_name : "file name",
			Link : "youtube.com",
		},
		
	}
	err := w.AddWorkShop()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	fmt.Println("test passed workshop added.")

}

func TestGetWorkshopRequest(t *testing.T) {
	w, err := GetWorkshopRequest()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	fmt.Println(w, err)

}
