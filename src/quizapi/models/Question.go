package models

type Question struct {

	ID string `json:"id,omitempty"`
	Title string `json:"question"`
	Options []string `json:"answers"`
	Answer int `json:"correct_answer"`
	TestID string `json:"test_id,omitempty"`

}


