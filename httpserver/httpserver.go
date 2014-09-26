package main

import (
	"encoding/json"
	"fmt"
	//"io"
	//"html"
	"io/ioutil"
	"log"
	"net/http"
)

type test_struct struct {
	name string
}

func main() {

	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/hellogo", HelloGo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello senthil , %q", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Hello senthil!")
}
func HelloGo(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	json.Unmarshal(keysBody, &keys)
	if err != nil {
		panic(err.Error())
	}
	log.Println(string(body))
	var t test_struct
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err.Error())
	}
	log.Println(t.Test)
	fmt.Fprintf(w, t.Test)
}
