package exeresponse

import (
	et "exectypes"
	"fmt"
)

func PostExecution_receiver(Response <-chan *et.ExecResponse) {
	for {
		rtrn := <-Response
		fmt.Println(rtrn.Response)
	}
}
