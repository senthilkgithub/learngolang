package exeresponse

import (
	"fmt"
	et "github.com/senthilkgithub/learngolang/exectypes"
)

func PostExecution_receiver(Response <-chan *et.ExecResponse) {
	for {
		rtrn := <-Response
		fmt.Println(rtrn.Response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
