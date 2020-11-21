package models

import (
	"domaniGoupdate/database"
	"fmt"

	_ "github.com/lib/pq" // Import solely for its initialization
)

type Domain struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	//Subdomains []Subdomain `json:"subdomains"`
	//SDNX int `json:"sdnx"`
	//TPCDX int `json:"tpcx"`
	//TSTDX int `json:"tstdx"`
}

/*
func (d *Domain) Update() (err error) {
	_, err = DB.GetCB().Bucket("test").Upsert(d.ID, &d, nil)
	return
}
*/
func (a *Domain) NewDomain() (id int, err error) {
	query := "INSERT INTO public.domain(id, title, image) VALUES ($1, $2, $3);"
	a.ID = NewUUID()

	// DB Query
	_, err = database.Get().Exec(query, a.ID, a.Title, a.Image)

	if err != nil {
		fmt.Println("err")
		id = -1
	}

	return
}
