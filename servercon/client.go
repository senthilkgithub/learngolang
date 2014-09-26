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
	//strEcho := "Halo"
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
	var yesOrNo string
	fmt.Println("Enter Some Text to send to Server")
	rtrn, err = fmt.Scanf("%s\n", &dataToSend)
	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.Write([]byte(dataToSend))
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

	fmt.Println("Do you want to continue again Y/N :")
	n, err := fmt.Scanf("%s\n", &yesOrNo)
	if err != nil || n != 1 {
		// handle invalid input
		fmt.Println(n, err)
		return
	}

	if yesOrNo == "y" || yesOrNo == "Y" {
		callServer()
	} else if yesOrNo != "n" || yesOrNo == "N" {
		fmt.Println("Wrong Code")
		callServer()
	}
	go conn.Close()

}
