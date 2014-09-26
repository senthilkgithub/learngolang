package main

import (
	//"fmt"
	"fmt"
	"net"
	"os"
)

var (
	dataToSend string
	rtrn       int
	err        error
	replyByte  []byte
)

func main() {
	callServer()

}
func callServer() {
	for i := 1; i < 100; i++ {
		servAddr := "localhost:3333"
		tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			println("Dial failed:", err.Error())
			os.Exit(1)
		}

		_, err = conn.Write([]byte("Auto Text"))
		if err != nil {
			fmt.Println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		//println("write to server = ", strEcho)

		replyByte = make([]byte, 1024)

		rtrn, err = conn.Read(replyByte)
		if err != nil {
			fmt.Println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		fmt.Println("reply from server=", string(replyByte))
		//go conn.CloseRead()

		conn.Close()
	}

}
