package cexec

import (
	et "exectypes"
	"fmt"
)

//recieve from structure
func C_Executer(request *et.ExecRequest, Response chan<- *et.ExecResponse) { //(cco <-chan *CompileCodeObj)
	fmt.Println("recieved ", request.Language, " for Execution")
	exeresult := new(et.ExecResponse)
	exeresult.Language = request.Language
	exeresult.Response = "c code has been executed successfully"
	exeresult.RequestCode = request.RequestCode
	Response <- exeresult
}
