package models

import (
	"fmt"

	"strconv"

	"domaniGoupdate/database"

	"github.com/lib/pq"

	_ "github.com/lib/pq" // Import solely for its initialization
)

type Test struct {
	ID   string `json:"id,omitempty"`
	Time int    `json:"time"`
	//SubdomainIDX int `json:"subdomain_id,omitempty"`
	//TopicIDX int `json:"topic_id,omitempty"`
	//Idx int `json:"idx,omitempty"`
	TpID string `json:"tp_id"`
}

type TestNest struct {
	ID       string     `json:"id,omitempty"`
	TpID     string     `json:"topic_id"`
	Time     int        `json:"time"`
	Question []Question `json:"questions"`
}

func (a *TestNest) NewTest() (id int, err error) {
	query := "INSERT INTO public.test(id, tp_id, time) VALUES ($1, $2, $3);"

	a.ID = NewUUID()
	// DB Query
	_, err = database.Get().Exec(query, a.ID, a.TpID, a.Time)

	if err != nil {
		fmt.Println("err")
		id = -1
	}
	query2 := "INSERT INTO public.question(id, title, options, answer, test_id) VALUES ($1, $2, $3, $4, $5);"

	var i int
	var model Question
	for i, model = range a.Question {
		model.ID = NewUUID() + strconv.Itoa(i)
		model.TestID = a.ID
		// DB Query
		_, err = database.Get().Exec(query2, model.ID, model.Title, pq.Array(model.Options), model.Answer, model.TestID)
		if err != nil {
			fmt.Println(err)
		}
	}

	return
}
