package models

import (
	"fmt"
	"testing"
)

func TestAddFiles(t *testing.T) {
	f := Files{

		Trainer_Id:   3,
		Role:         "role check ",
		OriginalName: "mahidul",
		Type:         "type checked",
		Name:         "moon",
		Link:         "facebook.com",
		Access:       "admin",
	}
	_,err := f.Add()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	fmt.Println("test passed files added.")

}

func TestGetFiles(t *testing.T) {
	f, err := GetAllFiles()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	fmt.Println(f, err)

}
