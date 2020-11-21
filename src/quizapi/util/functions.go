package util

import (

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetJsonFromUrlB(url string) ([]byte, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	jsonBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetJsonAndBind(url string, obj interface{}) error {
	resp, err := http.Get(url)
	if err!=nil{
		return err
	}
	defer resp.Body.Close()
	jsonBody, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonBody, &obj)
	if err != nil {
		return err
	}
	return nil
}

func GetJsonAndBindAuth(url string, token string, obj interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", token)
	//fmt.Println("2) Token from GETJSONBIND", token)
	// Prepare an HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	jsonBody, err := ioutil.ReadAll(resp.Body)
	//fmt.Println("3) GetJsonBind", string(jsonBody))
	err = json.Unmarshal(jsonBody, &obj)
	if err != nil {
		return err
	}
	return nil
}

func GetJsonFromUrlS(url string) (string, error) {
	jsonBytes, err := GetJsonFromUrlB(url)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func FileReader(filepath string) (string,error){
	byte, err := ioutil.ReadFile(filepath)
	file:=string(byte)

	return file,err
}


//package util
//
//import (
//	"encoding/json"
//	"io/ioutil"
//	"log"
//	"net/http"
//)
//
//// func GetJsonFromUrlB(url string) ([]byte, error) {
//// 	resp, err := http.Get(url)
//// 	defer resp.Body.Close()
//// 	jsonBody, err := ioutil.ReadAll(resp.Body)
//// 	if err != nil {
//// 		return nil, err
//// 	}
//// 	return jsonBody, nil
//// }
//
//// func GetJsonAndBind(url string, obj interface{}) error {
//// 	resp, err := http.Get(url)
//// 	defer resp.Body.Close()
//// 	jsonBody, err := ioutil.ReadAll(resp.Body)
//// 	err = json.Unmarshal(jsonBody, &obj)
//// 	if err != nil {
//// 		return err
//// 	}
//// 	return nil
//// }
//
//func GetJsonAndBindAuth(url string, token string, obj interface{}) error {
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return err
//	}
//	// Set appropriate headers
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Accept", "application/json")
//	req.Header.Set("Authorization", token)
//	//fmt.Println("2) Token from GETJSONBIND", token)
//	// Prepare an HTTP request
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Fatal(err.Error())
//		return err
//	}
//	jsonBody, err := ioutil.ReadAll(resp.Body)
//	//fmt.Println("3) GetJsonBind", string(jsonBody))
//	err = json.Unmarshal(jsonBody, &obj)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//// func GetJsonFromUrlS(url string) (string, error) {
//// 	jsonBytes, err := GetJsonFromUrlB(url)
//// 	if err != nil {
//// 		return "", err
//// 	}
//// 	return string(jsonBytes), nil
//// }
