package models

import (
	"domaniGoupdate/database"
	"fmt"

	_ "github.com/lib/pq" // Import solely for its initialization
)

type Topic struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	Questions int    `json:"questions"`
	Time      int    `json:"time"`
	//Tests []Test `json:"tests,omitempty"`
	//SubdomainIDX int `json:"subdomain_idx,omitempty"`
	//Idx int `json:"idx,omitempty"`
	SdID string `json:"sd_id"`
}

func (a *Topic) NewTopic() (id int, err error) {
	query := "INSERT INTO public.topic(id, title, image, questions, time, sd_id) VALUES ($1, $2, $3, $4, $5, $6);"
	a.ID = NewUUID()

	// DB Query
	_, err = database.Get().Exec(query, a.ID, a.Title, a.Image, a.Questions, a.Time, a.SdID)

	if err != nil {
		fmt.Println("err")
		id = -1
	}

	return
}
