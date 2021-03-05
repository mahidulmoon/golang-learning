package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	http.ListenAndServe("localhost:8000",nil)
}
