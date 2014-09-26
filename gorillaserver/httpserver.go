package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	et "github.com/senthilkgithub/learngolang/exectypes"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" //
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Myname struct {
	Name string `json:"name"`
}

type Comapny struct {
	ObjectId    int
	CompanyName string
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

	rtr.HandleFunc("/CompileCode", CompileCode)
	rtr.HandleFunc("/GetAllCompanyData", GetAllCompanyData)

	http.Handle("/", rtr)

	log.Println("Listening...", time.Now())

	http.ListenAndServe(":3000", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Hello " + name))
}

func GetAllCompanyData(w http.ResponseWriter, r *http.Request) {
	db := mysql.New("tcp", "", "devserver1:3306", "root", "lotus", "appserver_core")
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	res, err := db.Start("SELECT id,company_name FROM companys")
	if err != nil {
		panic(err)
	}
	//Print result to Response Writer
	Comapnys := make([]Comapny, 700)
	i := 0
	for {
		row, err := res.GetRow()
		if err != nil {
			panic(err)
		}
		if row == nil {
			break
		}
		Comapnys[i] = Comapny{row.Int(0), row.Str(1)}
		i++
	}
	db.Close()
	fmt.Println(time.Now())
	js, err := json.Marshal(Comapnys)
	w.Write(js)
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
	exeresult := et.ExecResponse{}
	exeresult.Language = request.Language
	exeresult.Response = "c++ code has been executed successfully"
	exeresult.RequestCode = request.RequestCode
	return exeresult
}
func Java_Executer(request *et.ExecRequest) et.ExecResponse {
	exeresult := et.ExecResponse{}
	exeresult.Language = request.Language
	exeresult.Response = "Java code has been executed successfully"
	exeresult.RequestCode = request.RequestCode
	return exeresult
}
