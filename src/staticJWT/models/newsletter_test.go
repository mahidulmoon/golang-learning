package models

import (
	"fmt"
	"testing"
)


func TestNewsletterAdd(t *testing.T) {
	ws := NewsLetter{
		
		ID:            0,
		Email:          "mahidulmoon.upskill@gmail.com",
		
	}
	err := ws.Add()
	if err != nil {
		fmt.Println("couldnot insert")
	}
}
