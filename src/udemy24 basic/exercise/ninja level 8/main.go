package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age int
}

type Person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}


func main(){
	u1 := user{
		Name: "mahidul",
		Age: 23,
	}
	u2 := user{
		Name: "moon",
		Age: 23,
	}
	u3 := user{
		Name: "meherafroz",
		Age: 23,
	}

	user := []user{u1,u2,u3} //making json format
	fmt.Println(user)

	bs,err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:",err)
	}
	//fmt.Println(bs)
	fmt.Println(string(bs)) //actual json format

	jsons := `[{"Name":"mahidul","Age":23},{"Name":"moon","Age":23},{"Name":"meherafroz","Age":23}]`

	var people []Person

	err = json.Unmarshal([]byte(jsons),&people)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(people)

}
