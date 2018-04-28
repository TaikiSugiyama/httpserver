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
	splitMethod := strings.Split(line, " ")
	return splitMethod[0]
}

func retrieveURIRequestLine(line string) (uri string) {
	splitMethod := strings.Split(line, " ")
	return splitMethod[1]
}

func retrieveVersionRequestLine(line string) (ver string) {
	splitMethod := strings.Split(line, " ")
	return splitMethod[2]
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
	var ver string
	var requestArray []string

	for _, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(line, -1) {
		requestArray = append(requestArray, v)
	}

	method = retrieveMethodRequestLine(requestArray[0])
	uri = retrieveURIRequestLine(requestArray[0])
	ver = retrieveVersionRequestLine(requestArray[0])

	//Write
	_, err = conn.Write([]byte(ver + " 200 OK\n"))
	checkError(err)
	_, err = conn.Write([]byte("Content-Type: text/html\n"))
	checkError(err)
	_, err = conn.Write([]byte("\n"))
	checkError(err)

	switch method {
	case "GET":
		filePath, _ := os.Getwd()
		filePath = filePath + uri
		if _, err := os.Stat(filePath); err != nil {
			fmt.Println("no such file: ", err)
			break
		}
		// 外部テンプレートファイルの読み込み
		idat, err := ioutil.ReadFile(filePath)
		checkError(err)
		_, err = conn.Write([]byte(idat))
		checkError(err)
	default:
		_, err = conn.Write([]byte("<h1>default</h1>\n"))
		checkError(err)
		fmt.Println(method)
	}

	fmt.Println("<<<< end")
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("fatal: error %s\n", err.Error())
		os.Exit(1)
	}
}
