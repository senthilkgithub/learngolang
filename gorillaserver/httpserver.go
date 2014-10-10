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
	CompilerRequest  et.ExecRequest
	CompilerResponse et.ExecResponse
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
	// CompilerRequest = make(chan et.ExecRequest)
	// CompilerResponse = make(chan et.ExecResponse)
	// go CompilerResponseWriter(CompilerResponse)
	log.Println("Listening...", time.Now())

	http.ListenAndServe(":3000", nil)
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
	switch jsonData.Language {
	case "c":
		CompilerResponse = cexec.C_Executer(jsonData)
	case "cpp":
		CompilerResponse = cppexec.CPlusPlus_Executer(jsonData)
	case "java":
		Java_Executer(rec, CompilerResponse)
	}
	go deletefolder.RemoveAllFiles(CompilerResponse.FolderPath)
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(CompilerResponse.Response)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(jsonResponse)
	}
}

/*func GetAllCompanyData(w http.ResponseWriter, r *http.Request) {
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
}*/
