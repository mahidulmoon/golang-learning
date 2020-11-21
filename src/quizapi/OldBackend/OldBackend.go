package OldBackend

import (
	"bytes"
	"domaniGoupdate/util"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type WorkshopData struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	Youtube_link string `json:"youtube_link"`
	Location     string `json:"location"`
	Duration     string `json:"duration"`
	Category     string `json:"category"`
	Fee          int    `json:"fee"`
	TS           string `json:"workshop_ts"`
}

type WorkshopsData struct {
	Data []WorkshopData `json:"data"`
}

type WorkshopsDataSingle struct {
	Data WorkshopData `json:"data"`
}

type OldBackend struct {
	Url string
}

var Instance = OldBackend{"https://backend.upskill.com.bd/api"}

func (o *OldBackend) GetWorshopsURL() string {
	return o.Url + "/workshops"
}

func (o *OldBackend) GetUserURL() string {
	return o.Url + "/user"
}

func (o *OldBackend) GetJoin(id string) string {
	return o.Url + "/workshop/" + id + "/join"
}

func (o *OldBackend) GetWorkshops() ([]WorkshopData, error) {
	workshopsData := WorkshopsData{}
	err := util.GetJsonAndBind(o.GetWorshopsURL(), &workshopsData)
	workshops := []WorkshopData{}
	for _, workshop := range workshopsData.Data {
		workshops = append(workshops, workshop)
	}
	workshopsData.Data = workshops
	if err != nil {
		return nil, err
	}
	return workshops, nil
}

func (o *OldBackend) GetWorkshopByID(id string) (WorkshopData, error) {
	workshopData := WorkshopsDataSingle{}
	err := util.GetJsonAndBind(o.GetWorshopsURL()+"/"+id, &workshopData)
	if err != nil {
		return WorkshopData{}, err
	}
	return workshopData.Data, nil
}

type UpskillUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (o *OldBackend) GetUserFromToken(token string) (UpskillUser, error) {
	user := UpskillUser{}
	err := util.GetJsonAndBindAuth(o.GetUserURL(), token, &user)
	return user, err
}

type RegisteredUser struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Token         string `json:"api_token"`
	Phone         string `json:"phone"`
	Status        string `json:"status"`
	VerifiedAt    string `json:"email_verified_at"`
	RememberToken string `json:"remember_token"`
	DeletedAt     string `json:"deleted_at"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	WorkshopID    int    `json:"workshop_id"`
	UserID        int    `json:"user_id"`
}

type RegisteredUsers struct {
	Data []RegisteredUser `json:"data"`
}

func (o *OldBackend) GetRegistrants(id string, token string) ([]RegisteredUser, error) {
	upskill := []RegisteredUser{
		{ID: 0,
			Name:          "Mustafizur R Khan",
			Email:         "",
			Password:      "",
			Token:         "",
			Phone:         "01819218449",
			Status:        "",
			VerifiedAt:    "",
			RememberToken: "",
			DeletedAt:     "",
			CreatedAt:     "",
			UpdatedAt:     "",
			WorkshopID:    0,
			UserID:        0},
		{
			ID:            0,
			Name:          "Ridwan Uz Zaman",
			Email:         "",
			Password:      "",
			Token:         "",
			Phone:         "01723978826",
			Status:        "",
			VerifiedAt:    "",
			RememberToken: "",
			DeletedAt:     "",
			CreatedAt:     "",
			UpdatedAt:     "",
			WorkshopID:    0,
			UserID:        0,
		},
		{
			ID:            0,
			Name:          "Sanan Hussain",
			Email:         "",
			Password:      "",
			Token:         "",
			Phone:         "01727234132",
			Status:        "",
			VerifiedAt:    "",
			RememberToken: "",
			DeletedAt:     "",
			CreatedAt:     "",
			UpdatedAt:     "",
			WorkshopID:    0,
			UserID:        0,
		},
		{
			ID:            0,
			Name:          "Aqib Hasan",
			Email:         "",
			Password:      "",
			Token:         "",
			Phone:         "01717272500",
			Status:        "",
			VerifiedAt:    "",
			RememberToken: "",
			DeletedAt:     "",
			CreatedAt:     "",
			UpdatedAt:     "",
			WorkshopID:    0,
			UserID:        0,
		},
		{
			ID:            0,
			Name:          "Kazi Lotus",
			Email:         "rafiul.startupdhaka@gmail.com",
			Password:      "",
			Token:         "",
			Phone:         "01707945494",
			Status:        "",
			VerifiedAt:    "",
			RememberToken: "",
			DeletedAt:     "",
			CreatedAt:     "",
			UpdatedAt:     "",
			WorkshopID:    0,
			UserID:        0,
		},
		{
			ID:            0,
			Name:          "Tasfia Jannat",
			Email:         "rafiul2111@gmail.com",
			Password:      "",
			Token:         "",
			Phone:         "01911781418",
			Status:        "",
			VerifiedAt:    "",
			RememberToken: "",
			DeletedAt:     "",
			CreatedAt:     "",
			UpdatedAt:     "",
			WorkshopID:    0,
			UserID:        0,
		},
		{
			ID:            0,
			Name:          "Rafiul",
			Email:         "rafiul2110@gmail.com",
			Password:      "",
			Token:         "",
			Phone:         "01723335418",
			Status:        "",
			VerifiedAt:    "",
			RememberToken: "",
			DeletedAt:     "",
			CreatedAt:     "",
			UpdatedAt:     "",
			WorkshopID:    0,
			UserID:        0,
		},
	}
	if id == "0" {
		return upskill, nil
	}

	data := RegisteredUsers{}
	url := "https://backend.upskill.com.bd/api/registrants/" + id
	//url := "http://localhost:3000/data"
	err := util.GetJsonAndBindAuth(url, token, &data)
	if err != nil {
		return nil, err
	}
	return data.Data, err
}

func (o *OldBackend) GetAllUsers(token string) ([]RegisteredUser, error) {
	data := RegisteredUsers{}
	url := "https://backend.upskill.com.bd/api/users"
	err := util.GetJsonAndBindAuth(url, token, &data)
	if err != nil {
		return nil, err
	}
	return data.Data, err
}

func (o *OldBackend) UserExistsByEmail(email string) bool {
	users, err := o.GetAllUsers("")
	if err != nil {
		return false
	}
	for _, v := range users {
		if v.Email == email {
			return true
		}
	}
	return false
}

type PasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (o *OldBackend) ChangePassword(email, password string) (resp string, err error) {
	requestBody := PasswordRequest{
		Email:    email,
		Password: password,
	}
	reqBody, err := json.Marshal(&requestBody)
	req, err := http.NewRequest("POST", o.Url+"/reset", bytes.NewBuffer(reqBody))
	//req.Header.Set("Authorization", "Basic aW5hbS5zdGFydHVwZGhha2FAZ21haWwuY29tOmluYW1iaGFp")
	req.Header.Set("Authorization", "gay")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	responseBody, err := ioutil.ReadAll(response.Body)
	return string(responseBody), err
}

// 27-01-2020

type SiteRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type SiteRegisterResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

func (o *OldBackend) Register(registrationInfo SiteRegister) (resp SiteRegisterResponse, err error) {
	reqBody, err := json.Marshal(&registrationInfo)
	req, err := http.NewRequest("POST", o.Url+"/register", bytes.NewBuffer(reqBody))
	//req.Header.Set("Authorization", "Basic aW5hbS5zdGFydHVwZGhha2FAZ21haWwuY29tOmluYW1iaGFp")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	responseBody, err := ioutil.ReadAll(response.Body)
	Response := SiteRegisterResponse{
		ID:    0,
		Token: "",
	}
	err = json.Unmarshal(responseBody, &Response)

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return Response, err
	} else {
		msg := string(responseBody)
		var err error
		if strings.Contains(msg, "Duplicate entry") {
			err = errors.New("Phone number is already registered!")
		} else if strings.Contains(msg, "The email has already been taken.") {
			err = errors.New("Email has already been registered in this site.")
		} else {
			err = errors.New("Something went wrong.")
		}
		return Response, err
	}
}

func (o *OldBackend) RegisterForWorkshopByUserTokenAndWorkshopID(userToken string, workshopId string) (resp map[string]string, err error) {
	req, err := http.NewRequest("GET", o.GetJoin(workshopId), nil)
	fmt.Println(req.URL)
	//req.Header.Set("Authorization", "Basic aW5hbS5zdGFydHVwZGhha2FAZ21haWwuY29tOmluYW1iaGFp")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", userToken)
	client := &http.Client{}
	response, err := client.Do(req)
	responseBody, err := ioutil.ReadAll(response.Body)
	Response := map[string]string{}
	err = json.Unmarshal(responseBody, &Response)
	return Response, nil
}
