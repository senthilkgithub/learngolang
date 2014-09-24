package javaexec

import (
	"fmt"
	et "github.com/senthilkgithub/learngolang/exectypes"
)

func Java_Executer(javaCodeRequest *et.ExecRequest, Response chan<- *et.ExecResponse) {
	fmt.Println("recieved ", javaCodeRequest.Language, " for Execution")
	exeresult := new(et.ExecResponse)
	exeresult.Language = javaCodeRequest.Language
	exeresult.Response = "Java code has been executed successfully"
	exeresult.RequestCode = javaCodeRequest.RequestCode
	Response <- exeresult

}
