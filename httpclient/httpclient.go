package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type jsonStructure struct {
	IntParam1, IntParam2       int
	StringParam1, StringParam2 string
}
type CandidateStructure struct {
	Name, Dob, Address, City, State, Degree, Gender, Occupation string
	Age                                                         int
}

var (
	fileName          string
	jsonObj           jsonStructure
	jsonReq           []byte
	candidateObj      CandidateStructure
	candidateReq      []byte
	loggerChannel         = make(chan string)
	routineCompleted      = make(chan int, 5)
	completedchannels int = 0
	helloChan             = make(chan int)
	pyramidChan           = make(chan int)
	JsonChan              = make(chan int)
	helloDbChan           = make(chan int)
	CreateCandChan        = make(chan int)
	ResponseTime      string
)

func main() {
	f, err := ioutil.TempFile(os.TempDir(), "compareservicelog")

	if err != nil {
		fmt.Println(err.Error())
	}
	fileName = f.Name()

	fmt.Println(fileName)
	jsonObj = jsonStructure{}
	jsonObj.IntParam1 = 1
	jsonObj.IntParam2 = 2
	jsonObj.StringParam1 = "Test String1"
	jsonObj.StringParam2 = "Test String2"
	jsonReq, err = json.Marshal(jsonObj)
	if err != nil {
		fmt.Println(err.Error())
	}
	candidateObj = CandidateStructure{}
	candidateObj.Address = "India Gate"
	candidateObj.Age = 45
	candidateObj.City = "New Delhi"
	candidateObj.Degree = "B Tech"
	candidateObj.Dob = "1970-01-01 00:00:00"
	candidateObj.Gender = "UnSpecified"
	candidateObj.Name = "Krishna"
	candidateObj.Occupation = "Minister"
	candidateObj.State = "Delhi"
	candidateReq, err = json.Marshal(candidateObj)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("All Request Begins At", time.Now())
	go WriteIntoTempFile(loggerChannel)
	go CompleteMyOPerations(1)
	go CallHelloMethod()
	go CallPyramidConstructor()
	go CallHelloJson()
	go CallHelloDb()
	go CallCreateCandidate()
	go HelloMethodFunction(helloChan)
	go PyramidConstructorFunction(pyramidChan)
	go HelloJsonFunction(JsonChan)
	go HelloDbFunction(helloDbChan)
	go CreateCandidateFunction(CreateCandChan)
	fmt.Println("Please wait")
	getChar()

}

func CompleteMyOPerations(id int) {
	for {
		_ = <-routineCompleted
		completedchannels++
		if completedchannels == 5 {
			f, err := os.OpenFile(fileName, os.O_APPEND, 0666)
			if err != nil {
				fmt.Println(err.Error())
			}
			_, _ = f.WriteString(ResponseTime)

			f.Close()
			fmt.Println("All Requests completed At", time.Now())
		}
	}
}

func WriteIntoTempFile(logger <-chan string) {
	for {
		out := <-loggerChannel
		ResponseTime = ResponseTime + out + "\n"

	}
}

func CallHelloMethod() {
	for i := 1; i <= 10000; i++ {
		helloChan <- i
	}
	routineCompleted <- 1

}
func HelloMethodFunction(helloChan <-chan int) {
	for {
		_ = <-helloChan
		requestStart := time.Now()
		resp, err := http.Get("http://senthilkumar:3000/hello")
		if err != nil {
			fmt.Println(err.Error())
		}
		resp.Body.Close()
		requestEnd := time.Now()
		loggerChannel <- "hello," + requestEnd.Sub(requestStart).String()
	}
}
func CallPyramidConstructor() {

	for i := 1; i <= 10000; i++ {
		pyramidChan <- i
	}
	routineCompleted <- 2

}
func PyramidConstructorFunction(pyramidChan <-chan int) {
	for {
		_ = <-pyramidChan
		requestStart := time.Now()
		resp, _ := http.Get("http://senthilkumar:3000/StarPyramidGenerator/15")
		resp.Body.Close()
		requestEnd := time.Now()
		loggerChannel <- "PyramidConstructor," + requestEnd.Sub(requestStart).String()
	}
}

func CallHelloJson() {

	go func() {
		for i := 1; i <= 10000; i++ {
			JsonChan <- i
		}
		routineCompleted <- 3
	}()
}
func HelloJsonFunction(JsonChan <-chan int) {
	for {
		_ = <-JsonChan
		requestStart := time.Now()
		req, err := http.NewRequest("POST", "http://senthilkumar:3000/helloJson", bytes.NewBuffer(jsonReq))
		if err != nil {
			fmt.Println(err.Error())
		}
		req.Header.Add("User-Agent", "myClient")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		resp.Body.Close()
		requestEnd := time.Now()
		loggerChannel <- "HelloJson," + requestEnd.Sub(requestStart).String()
	}
}

func CallHelloDb() {
	go func() {
		for i := 1; i <= 10000; i++ {
			helloDbChan <- i
		}
		routineCompleted <- 4
	}()
}
func HelloDbFunction(helloDbChan <-chan int) {
	for {
		_ = <-helloDbChan
		requestStart := time.Now()
		req, err := http.NewRequest("POST", "http://senthilkumar:3000/helloDb", bytes.NewBuffer(jsonReq))
		if err != nil {
			fmt.Println(err.Error())
		}
		req.Header.Add("User-Agent", "myClient")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		resp.Body.Close()
		requestEnd := time.Now()
		//fmt.Println("helloDb," + requestEnd.Sub(requestStart).String())
		loggerChannel <- "helloDb," + requestEnd.Sub(requestStart).String()
	}
}

func CallCreateCandidate() {
	go func() {
		for i := 1; i <= 10000; i++ {
			CreateCandChan <- i
		}
		routineCompleted <- 5

	}()
}

func CreateCandidateFunction(CreateCandChan <-chan int) {
	for {
		_ = <-CreateCandChan
		requestStart := time.Now()
		req, err := http.NewRequest("POST", "http://senthilkumar:3000/CreateCandidate", bytes.NewBuffer(candidateReq))
		if err != nil {
			fmt.Println(err.Error())
		}
		req.Header.Add("User-Agent", "myClient")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		resp.Body.Close()
		requestEnd := time.Now()
		loggerChannel <- "CreateCandidate," + requestEnd.Sub(requestStart).String()
	}
}

func getChar() {
	var inputKey string
	count, _ := fmt.Scanf("%s\n", &inputKey)
	if count != 1 {
		fmt.Println(inputKey)
	}
}
