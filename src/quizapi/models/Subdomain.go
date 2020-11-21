package models

import (
	"fmt"

	"domaniGoupdate/database"

	_ "github.com/lib/pq" // Import solely for its initialization
)

type Subdomain struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
	Image string `json:"image"`
	//Topics []Topic `json:"topics,omitempty"`
	//Idx int `json:"idx,omitempty"`
	DID string `json:"did"`
}

func (a *Subdomain) NewSubdomain() (id int, err error) {
	query := "INSERT INTO public.subdomain(id, title, image, did) VALUES ($1, $2, $3, $4);"
	a.ID = NewUUID()

	// DB Query
	_, err = database.Get().Exec(query, a.ID, a.Title, a.Image, a.DID)

	if err != nil {
		fmt.Println("err")
		id = -1

	}

	return
}
