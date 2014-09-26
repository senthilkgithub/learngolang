package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var inputKey string

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	// http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	count := 0
	count, _ = fmt.Scanf("%s", &inputKey)
	if count <= 0 {
		fmt.Scanf("%s", &inputKey)
	}

}
