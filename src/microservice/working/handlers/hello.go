package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter,r *http.Response){
	h.l.Println("Hello world")
	d,err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(rw,"Oops",http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw,"Data %s",d)
}