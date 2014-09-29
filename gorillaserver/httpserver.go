//Compiler web server using Gorilla Mux
package main

//Import the required packages,
import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/senthilkgithub/learngolang/cexec"
	"github.com/senthilkgithub/learngolang/cppexec"
	"github.com/senthilkgithub/learngolang/deletefolder"
	et "github.com/senthilkgithub/learngolang/exectypes"
	//"github.com/ziutek/mymysql/mysql"
	"fmt"
	_ "github.com/ziutek/mymysql/native" //
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/* Create two channel variable which will carry input request ,and output response
Name it as CompilerRequest,CompilerResponse
*/
var (
	CompilerRequest  chan et.ExecRequest
	CompilerResponse chan et.ExecResponse
)

/*
Main is the entry point of the execution
Creates Router and assigns Header parameters
This http app listen at port number 3000
For here application  supports Get and Post method for Getting & Updating operation
The router is handling Compile Code method of type post,
The client has to send
{
	Language    (c,++,java) (Text type)
	Code         (Source code as text as it is) (Text Type)
	CommandLineArgs  (param1,param2,...) (Text Type)
}
*/
func main() {
	rtr := mux.NewRouter()
	rtr.Headers("Content-Type", "application/json",
		"X-Requested-With", "XMLHttpRequest")
	rtr.Methods("GET,POST")
	rtr.HandleFunc("/CompileCode", CompileCode)

	http.Handle("/", rtr)
	CompilerRequest = make(chan et.ExecRequest)
	CompilerResponse = make(chan et.ExecResponse)
	go CompilerResponseWriter(CompilerResponse)
	log.Println("Listening...", time.Now())

	http.ListenAndServe(":3000", nil)
}

// func Profile(w http.ResponseWriter, r *http.Request) {
// 	MynameObj := Myname{"senthil", w}
// 	go RequestResponder(Req)
// 	requestSender(Req, MynameObj)

// }

/*requestSender is simple metheod receive request and upstream request to the channel CompilerRequest*/
func requestSender(Req chan et.ExecRequest, reqObj et.ExecRequest) { //req chan<- Myname,
	CompilerRequest <- reqObj
}

/*RequestResponder is goRoutine metheod receive requests  in channel forever and calls corresponding executer,
request and upstream request to the channel CompilerRequest*/

func RequestResponder(CompilerRequest chan et.ExecRequest, CompilerResponse chan et.ExecResponse) {
	for {
		rec := <-CompilerRequest
		switch rec.Language {
		case "c":
			cexec.C_Executer(rec, CompilerResponse)
		case "c++":
			cppexec.CPlusPlus_Executer(rec, CompilerResponse)
			// case "java":
			// 	Java_Executer(rec, CompilerResponse)
		}
	}
}

/*CompilerResponseWriter is goRoutine metheod receive final Response in channel forever and write in to the corresponding response writer*/
func CompilerResponseWriter(CompilerResponse <-chan et.ExecResponse) {
	for {
		resp := <-CompilerResponse
		go deletefolder.RemoveAllFiles(resp.FolderPath)
		resp.RespWriter.Header().Set("Content-Type", "application/json")
		// jsonResponse, err := json.Marshal(resp.Response)
		// if err != nil {
		// 	resp.RespWriter.Write([]byte(err.Error()))
		// } else {
		fmt.Println(resp.Response)
		resp.RespWriter.Write([]byte(resp.Response))
		//		}

	}
}

/*mux serve http function for CompileCode */
func CompileCode(w http.ResponseWriter, req *http.Request) {

	jsonRequestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData := et.ExecRequest{}
	err = json.Unmarshal(jsonRequestBody, &jsonData)
	req.Body.Close()
	jsonData.RespWriter = w
	go RequestResponder(CompilerRequest, CompilerResponse)
	requestSender(CompilerRequest, jsonData)
}

/*func CPlusPLus_Executer(request *et.ExecRequest) et.ExecResponse { //(cco <-chan *CompileCodeObj)
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
}*/

//func GetAllCompanyData(w http.ResponseWriter, r *http.Request) {
// 	db := mysql.New("tcp", "", "devserver1:3306", "root", "lotus", "appserver_core")
// 	err := db.Connect()
// 	if err != nil {
// 		panic(err)
// 	}
// 	res, err := db.Start("SELECT id,company_name FROM companys")
// 	if err != nil {
// 		panic(err)
// 	}
// 	//Print result to Response Writer
// 	Comapnys := make([]Comapny, 700)
// 	i := 0
// 	for {
// 		row, err := res.GetRow()
// 		if err != nil {
// 			panic(err)
// 		}
// 		if row == nil {
// 			break
// 		}
// 		Comapnys[i] = Comapny{row.Int(0), row.Str(1)}
// 		i++
// 	}
// 	db.Close()
// 	fmt.Println(time.Now())
// 	js, err := json.Marshal(Comapnys)
// 	w.Write(js)
// }
