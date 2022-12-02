package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "sync"
	"time"

	"github.com/gorilla/mux"
)

var mapresponse = make(map[string]bool)

func main() {
	// fmt.Print("this just test")
	r := mux.NewRouter()
	r.HandleFunc("/apies/", handlerFunc).Methods("POST")
	r.HandleFunc("/test", handlerFuncGet).Methods("GET")
	go CallingParent()
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
func handlerFuncGet(rw http.ResponseWriter, re *http.Request) {
	fmt.Println("just test")
}
func CallingParent() {
	for {
		u := []string{"http://localhost:3000", "https://www.google.com/", "https://www.facebook.com/", "https://www.face.com/"}
		// ctx := make(chan int)
		Calling(u, mapresponse)
		fmt.Println(mapresponse)
		fmt.Println(" updated data")
		time.Sleep(time.Minute * 1)
		fmt.Println("wake up again->")

	}
}
func handlerFunc(rw http.ResponseWriter, re *http.Request) {
	fmt.Print("api call ->->->")
	re.ParseForm()
	maped := re.Form
	resmap := map[string]bool{}
	if len(maped) == 0 {
		json.NewEncoder(rw).Encode(mapresponse)
		return
	}
	for key := range maped {
		value := maped[key]
		resmap[value[0]] = mapresponse[value[0]]
	}
	json.NewEncoder(rw).Encode(resmap)
}

func Calling(linklist []string, mapres map[string]bool) {
	objectStatusChecker := httpChecker{}
	for _, val := range linklist {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*100))
		// wg.Add(1)
		status, err := objectStatusChecker.Check(ctx, val)
		mapres[val] = status
		if err != nil {
			fmt.Println(err)
		}
		defer cancel()
	}

}

type StatusChecker interface {
	Check(ctx context.Context, link string) (status bool, err error)
}

type httpChecker struct {
}

func (h httpChecker) Check(ctx context.Context, link string) (status bool,
	err error) {
	// your implementation to check status using HTTP call
	client := http.Client{}
	fmt.Println("this is call from check")
	r, er := http.NewRequest("GET", link, nil)
	if er != nil {
		return false, er
	}
	resp, err := client.Do(r)
	if err != nil {
		return false, er
	}

	return resp.Status == "200 OK", err
}

func Requester(link string) string {
	client := http.Client{}
	r, er := http.NewRequest("GET", link, nil)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println("at line 54", r, link)
	resp, err := client.Do(r)
	fmt.Println(resp, "response")
	if err != nil {
		fmt.Println(er)
	}
	fmt.Println(resp.Status, "this is status ")
	return resp.Status

}
