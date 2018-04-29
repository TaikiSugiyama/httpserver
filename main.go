package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"

	"./handler"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	err = handleListener(listener)
	checkError(err)
}

func handleListener(listener *net.TCPListener) error {
	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			if ne, ok := err.(net.Error); ok {
				if ne.Temporary() {
					continue
				}
			}
			return err
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	defer conn.Close()
	fmt.Println("start >>>")

	lineBuf := make([]byte, 1024)
	lineLen, err := conn.Read(lineBuf)
	checkError(err)

	line := string(lineBuf[:lineLen])

	// read確認用
	fmt.Println(line)

	var method string
	var uri string
	var contentType string
	var ver string
	var requestArray []string
	typeMap := map[string]string{
		"html": "text/html",
		"css":  "text/css",
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"gif":  "image/gif",
	}

	for _, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(line, -1) {
		requestArray = append(requestArray, v)
	}

	method = handler.RetrieveMethodRequestLine(requestArray[0])
	uri = handler.RetrieveURIRequestLine(requestArray[0])
	contentType = typeMap[handler.RetrieveTypeRequestLine(uri)]
	ver = handler.RetrieveVersionRequestLine(requestArray[0])
	fmt.Println(contentType)

	returnResponse(conn, method, uri, ver, contentType)

	fmt.Println("<<<< end")
}

func returnResponse(conn *net.TCPConn, method string, uri string, ver string, contentType string) {

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
