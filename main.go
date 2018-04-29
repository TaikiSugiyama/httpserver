package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"strings"
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

func retrieveMethodRequestLine(line string) (method string) {
	split := strings.Split(line, " ")
	return split[0]
}

func retrieveURIRequestLine(line string) (uri string) {
	split := strings.Split(line, " ")
	return split[1]
}

func retrieveVersionRequestLine(line string) (ver string) {
	split := strings.Split(line, " ")
	return split[2]
}

func retrieveTypeRequestLine(uri string) (extension string) {
	start := strings.LastIndex(uri, ".")
	r := []rune(uri)
	return string(r[start:])
}

// func retrieveContentLength(line string) {
// 	regexpContentLen := regexp.MustCompile(`Content-Length`)
// 	if regexpContentLen.MatchString(line) {
// 		contentLength := strings.Split(line, ":\\s")[1]
// 		fmt.Println("content get :", contentLength)
// 		return contentLength
// 	}
// }

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

	for _, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(line, -1) {
		requestArray = append(requestArray, v)
	}

	method = retrieveMethodRequestLine(requestArray[0])
	uri = retrieveURIRequestLine(requestArray[0])
	contentType = retrieveTypeRequestLine(uri)
	ver = retrieveVersionRequestLine(requestArray[0])

	returnResponse(conn, method, uri, ver, contentType)

	fmt.Println("<<<< end")
}

func returnResponse(conn *net.TCPConn, method string, uri string, ver string, contentType string) {
	switch method {
	case "GET":
		filePath, _ := os.Getwd()
		filePath = filePath + uri
		if _, err := os.Stat(filePath); err != nil {
			fmt.Println("no such file: ", err)
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
	default:
		_, err := conn.Write([]byte("<h1>default</h1>\n"))
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
