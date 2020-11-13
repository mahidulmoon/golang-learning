package db

import (
	"fmt"
	"testing"
)

func TestIsDBAlive(t *testing.T) {
	db := GetDB()
	fmt.Println(db)
}

