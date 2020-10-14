package models

import (
	"fmt"
	"testing"
)

func TestAddFiles(t *testing.T) {
	f := Files{

		Trainer_Id:   31,
		Role:         "role check again",
		OriginalName: "mahidul",
		Type:         "typechecked",
		Name:         "moon",
		Link:         "youtube.com",
		Access:       "trainer",
	}
	err := f.Add()
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
