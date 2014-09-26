package main

import (
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	et "github.com/senthilkgithub/learngolang/exectypes"
	"io/ioutil"
	"log"
	"net/http"
)

type Myname struct {
	Name string `json:"name"`
}

var (
	CompilerRequest  et.ExecRequest
	CompilerResponse et.ExecResponse
)

func main() {
	rtr := mux.NewRouter()
	rtr.Headers("Content-Type", "application/json",
		"X-Requested-With", "XMLHttpRequest")
	rtr.Methods("GET,POST")

	rtr.HandleFunc("/GetJsonReq", GetJsonReq)
	rtr.HandleFunc("/CompileCode", CompileCode)

	http.Handle("/", rtr)

	log.Println("Listening...")

	// CompilerRequest := make(chan *et.ExecRequest)
	// CompilerResponse := make(chan *et.ExecResponse)
	// JsonRequest := make(chan *http.Request)
	// JsonResponse := make(chan *http.Response)
	// go RequestReceiver(JsonRequest, CompilerRequest)
	// go Executer(Request, Response)
	// go res.ResponseReceiver(Response)

	http.ListenAndServe(":3000", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func CompileCode(w http.ResponseWriter, req *http.Request) {

	jsonRequestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData := et.ExecRequest{}
	err = json.Unmarshal(jsonRequestBody, &jsonData)
	req.Body.Close()
	compilerResponse := Executer(&jsonData)
	jsonResponse, err2 := json.Marshal(compilerResponse)
	if err2 != nil {
		log.Panic(err2.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func Executer(codeStruct *et.ExecRequest) et.ExecResponse {
	switch codeStruct.Language {
	case "c":
		CompilerResponse = C_Executer(codeStruct)
	case "c++":
		CompilerResponse = CPlusPLus_Executer(codeStruct)
	case "java":
		CompilerResponse = Java_Executer(codeStruct)
	}
	return CompilerResponse
}
func C_Executer(request *et.ExecRequest) et.ExecResponse { //(cco <-chan *CompileCodeObj)
	exeresult := et.ExecResponse{}
	exeresult.Language = request.Language
	exeresult.Response = "c code has been executed successfully"
	exeresult.RequestCode = request.RequestCode
	return exeresult
}
func CPlusPLus_Executer(request *et.ExecRequest) et.ExecResponse { //(cco <-chan *CompileCodeObj)
	fmt.Println("recieved ", request.Language, " for Execution")
	exeresult := et.ExecResponse{}
	exeresult.Language = request.Language
	exeresult.Response = "c++ code has been executed successfully"
	exeresult.RequestCode = request.RequestCode
	return exeresult
}
func Java_Executer(request *et.ExecRequest) et.ExecResponse {
	fmt.Println("recieved ", javaCodeRequest.Language, " for Execution")
	exeresult := et.ExecResponse{}
	exeresult.Language = javaCodeRequest.Language
	exeresult.Response = "Java code has been executed successfully"
	exeresult.RequestCode = javaCodeRequest.RequestCode
	return exeresult
}
