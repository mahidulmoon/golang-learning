package models

import (
	"domaniGoupdate/database"
	"fmt"

	_ "github.com/lib/pq" // Import solely for its initialization
)

type TestSession struct {
	ID  string `json:"id"`
	UID int    `json:"uid"`
	//new
	Test_ID string `json:"test_id"`
}

func (a *TestSession) NewSession() (id int, err error) {
	query := "INSERT INTO public.testsession(id, uid, test_id) VALUES ($1, $2, $3);"
	a.ID = NewUUID()

	// DB Query
	_, err = database.Get().Exec(query, a.ID, a.UID, a.Test_ID)

	if err != nil {
		fmt.Println("err")
		id = -1
	}

	return
}
