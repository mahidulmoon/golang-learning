package models

import (
	"domaniGoupdate/database"
	"fmt"
)

type SubmissionModel struct {
	Correct float64 `json:"correct"`
	Total   float64 `json:"total"`
	TestId  string  `json:"test_id"`
	Token   string  `json:"token"`
}

func SubmitResult(userID int, submission SubmissionModel) error {
	query := "INSERT INTO public.testsession(id, uid, test_id, score) VALUES ($1, $2, $3, $4);"
	score := int(float64(submission.Correct) / float64(submission.Total) * 100)
	// DB Query
	_, err := database.Get().Exec(query, NewUUID(), userID, submission.TestId, score)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func GetStats(userID int) (float32, error) {
	query := "SELECT score from public.testsession WHERE uid = $1;"
	// DB Query
	rows, err := database.Get().Query(query, userID)

	if err != nil {
		return 0, err
	}
	sum := 0
	count := 0
	for rows.Next() {
		var num int
		rows.Scan(&num)
		sum += num
		count++
	}
	average := sum / count
	return float32(average), nil

}
