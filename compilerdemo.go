package main

import (
	"fmt"
	cpp "github.com/senthilkgithub/learngolang/c++exec"
	c "github.com/senthilkgithub/learngolang/cexec"
	et "github.com/senthilkgithub/learngolang/exectypes"
	res "github.com/senthilkgithub/learngolang/exeresponse"
	java "github.com/senthilkgithub/learngolang/javaexec"
)

var (
	selection int
	yesOrNo   string
)

func main() {
	Request := make(chan *et.ExecRequest)
	Response := make(chan *et.ExecResponse)
	go Executer(Request, Response)
	go res.PostExecution_receiver(Response)
	requestCapture(Request)
}

func requestCapture(Request chan<- *et.ExecRequest) {
	// Request := make(chan *et.ExecRequest)
	// go Executer(Request)
	availableOperations := [4]string{
		"1.C",
		"2.C++",
		"3.Java",
		"4.GoLang",
	}
	for _, value := range availableOperations {
		fmt.Println(value)
	}

	fmt.Println("Select your Compilation :")
	n, err := fmt.Scanf("%d\n", &selection)
	if err != nil || n != 1 {
		// handle invalid input
		fmt.Println(n, err)
		return
	}

	if selection > 0 {
		switch selection {
		case 1:
			go requestSender(Request, "c")
		case 2:
			go requestSender(Request, "c++")
		case 3:
			go requestSender(Request, "java")
		case 4:
			go requestSender(Request, "golang")

		}
		fmt.Println("Do you want to continue again Y/N :")
		n, err := fmt.Scanf("%s\n", &yesOrNo)
		if err != nil || n != 1 {
			// handle invalid input
			fmt.Println(n, err)
			return
		}

		if yesOrNo == "y" || yesOrNo == "Y" {
			requestCapture(Request)
		} else if yesOrNo != "n" || yesOrNo == "N" {
			fmt.Println("Wrong Code")
			requestCapture(Request)
		}

	}
	return
}

func requestSender(request chan<- *et.ExecRequest, lang string) { //requestSender(request chan<- string) {
	if lang != "" {
		fmt.Println("Sent ", lang, " for execution")
		compileobj := new(et.ExecRequest)
		compileobj.Language = lang
		compileobj.IsCompileOperation = false
		request <- compileobj
	}

}

func Executer(request <-chan *et.ExecRequest, Response chan<- *et.ExecResponse) {
	// Response := make(chan *et.ExecResponse)
	// go res.PostExecution_receiver(Response)
	for {
		codeStruct := <-request
		switch codeStruct.Language {
		case "c":
			go c.C_Executer(codeStruct, Response)
		case "c++":
			go cpp.CPlusPLus_Executer(codeStruct, Response)
		case "java":
			go java.Java_Executer(codeStruct, Response)
		}
	}
}
