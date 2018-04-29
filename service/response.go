package service

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func ReturnResponse(conn *net.TCPConn, method string, uri string, ver string, contentType string) {
	pwd, _ := os.Getwd()
	filePath := pwd + uri
	switch method {
	case "GET":
		if _, err := os.Stat(filePath); err != nil {
			_, err := conn.Write([]byte(ver + " 404 Not Found\n"))
			checkError(err)

			// header
			_, err = conn.Write([]byte("Content-Type: " + contentType + "\n"))
			checkError(err)
			_, err = conn.Write([]byte("\n"))
			checkError(err)

			// body
			body, err := ioutil.ReadFile(pwd + "/notfound.html")
			checkError(err)
			_, err = conn.Write([]byte(body))
			checkError(err)
			break
		}

		// status line
		_, err := conn.Write([]byte(ver + " 200 OK\n"))
		checkError(err)

		// header
		_, err = conn.Write([]byte("Content-Type: " + contentType + "\n"))
		checkError(err)
		_, err = conn.Write([]byte("\n"))
		checkError(err)

		// body
		body, err := ioutil.ReadFile(filePath)
		checkError(err)
		_, err = conn.Write([]byte(body))
		checkError(err)
	case "HEAD":
		if _, err := os.Stat(filePath); err != nil {
			_, err := conn.Write([]byte(ver + " 404 Not Found\n"))
			checkError(err)
			break
		}
		// status line
		_, err := conn.Write([]byte(ver + " 200 OK\n"))
		checkError(err)
		_, err = conn.Write([]byte("\n"))
		checkError(err)
		body, err := ioutil.ReadFile(filePath)
		checkError(err)
		_, err = conn.Write([]byte(body))
		checkError(err)
	default:
		_, err := conn.Write([]byte(ver + " 405 Method Not Allowed\n"))
		checkError(err)
		fmt.Println(method)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("fatal: error %s\n", err.Error())
		os.Exit(1)
	}
}
