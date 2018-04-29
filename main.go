package main

import (
	"fmt"
	"httpserver/handler"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	err = handler.HandleListener(listener)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("fatal: error %s\n", err.Error())
		os.Exit(1)
	}
}
