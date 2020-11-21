package models

import (
	"fmt"
	"strconv"
	"time"

	"domaniGoupdate/database"

	"github.com/lib/pq"

	_ "github.com/lib/pq" // Import solely for its initialization
)

// Helper functions
func NewUUID() string {
	//id, err := uuid.NewUUID()
	//if err != nil {
	//	return string(time.Now().UnixNano())
	//}
	//return id.String()
	return strconv.FormatInt(time.Now().Unix(), 10)
}

//-----------------WE NO LONGER HAVE A ROOT
/*
func GetRoot() (domains []Domain, err error) {
	query := "SELECT test.* FROM test"
	rows, err := db.GetCB().Cluster().Query(query, nil)

	for rows.Next() {
		var model Domain
		err = rows.Row(&model)
		domains = append(domains, model)
	}

	return
}
*/

func GetDomains() (domains []Domain, err error) {
	fmt.Println("getting domains")
	query := "SELECT * from domain"
	rows, err := database.Get().Query(query)
	if err != nil {
		println(err.Error)
	}
	for rows.Next() {
		var model Domain
		err = rows.Scan(&model.ID, &model.Title, &model.Image)
		domains = append(domains, model)
	}
	return
}

func GetSubdomainsByDomainID(domain_id string) (subdomains []Subdomain, err error) {
	fmt.Println("getting sub_domains")
	query := fmt.Sprintf("SELECT * from subdomain where did='%s'", domain_id)
	fmt.Println(query)
	rows, err := database.Get().Query(query)
	if err != nil {
		println(err.Error)
	}
	for rows.Next() {
		var model Subdomain
		err = rows.Scan(&model.ID, &model.Title, &model.Image, &model.DID)
		subdomains = append(subdomains, model)
	}
	return
}

func GetTopicsBySubdomainID(subdomain_id string) (topics []Topic, err error) {
	fmt.Println("getting_topics")
	query := fmt.Sprintf("SELECT * from topic where sd_id='%s'", subdomain_id)
	fmt.Println(query)
	rows, err := database.Get().Query(query)
	if err != nil {
		println(err.Error)
	}
	for rows.Next() {
		var model Topic
		err = rows.Scan(&model.ID, &model.Title, &model.Image, &model.Questions, &model.Time, &model.SdID)
		topics = append(topics, model)
	}
	return
}

func GetTestsByTopicID(topic_id string) (nests []TestNest, err error) {
	fmt.Println("In function")
	fmt.Println("getting tests for topic " + topic_id)
	query := "SELECT * from public.test WHERE tp_id=$1"
	rows, err := database.Get().Query(query, topic_id)
	if err != nil {
		fmt.Println("couldnt Run test")
	}

	for rows.Next() {
		var model Test
		var nestmodel TestNest
		err = rows.Scan(&model.ID, &model.Time, &model.TpID)
		items, err := GetQuestionByTestID(model.ID)
		if err != nil {
			fmt.Println("Nesting error")
		}
		nestmodel.Question = items
		nestmodel.ID = model.ID
		nestmodel.TpID = model.TpID
		nestmodel.Time = model.Time

		nests = append(nests, nestmodel)

	}

	return
}

//----------------Not used by handlers, Directly used by GetTestsByTopicId
func GetQuestionByTestID(test_id string) (models []Question, err error) {
	fmt.Println("getting questions for test " + test_id)
	query := "SELECT * from public.question WHERE test_id=$1"
	rows, err := database.Get().Query(query, test_id)
	if err != nil {
		fmt.Println("Question query error")
		fmt.Println(err)
	}
	for rows.Next() {
		var model Question
		err := rows.Scan(&model.ID, &model.Title, pq.Array(&model.Options), &model.Answer, &model.TestID)
		if err != nil {
			fmt.Println("error in model")
			return models, err
		}
		models = append(models, model)
	}
	fmt.Println(len(models))
	return
}

//---------------------ADD FUNCTIONS MOVED TO MODELS
