package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"text/template"
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

	line := lineBuf[:lineLen]
	//	fmt.Println(string(line))

	regexpContentLen := regexp.MustCompile(`Content-Length`)
	//リクエストメソッド　取得する　とりあえずは
	requestMethod := regexp.MustCompile(`GET`)
	var method string
	var path string

	for _, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(line), -1) {
		fmt.Println(string(v))
		if requestMethod.MatchString(v) {
			splitMethod := strings.Split(v, " ")
			method = splitMethod[0]
			path = splitMethod[1]
		}
		if regexpContentLen.MatchString(v) {
			contentLength := strings.Split(v, ":\\s")[1]
			fmt.Println("content get :", contentLength)
		}
	}

	//Write
	// response, err := conn.Write([]byte(line))
	_, err = conn.Write([]byte("HTTP/1.0 200 OK\n"))
	checkError(err)
	_, err = conn.Write([]byte("Content-Type: text/html\n"))
	checkError(err)
	_, err = conn.Write([]byte("\n"))
	checkError(err)

	fmt.Println("method", method)

	switch method {
	case "GET":
		filePath, _ := os.Getwd()
		filePath = filePath + path
		if _, err := os.Stat(filePath); err != nil {
			fmt.Println("no such file: ", err)
			break
		}
		// 外部テンプレートファイルの読み込み
		var t = template.Must(template.ParseFiles(filePath))
		// テンプレート出力
		if err := t.Execute(os.Stdout, nil); err != nil {
			fmt.Println(err.Error())
		}
	default:
		_, err = conn.Write([]byte("<h1>Hello World!!</h1>\n"))
		checkError(err)
		fmt.Println(method)
	}
	//	_, err = conn.Write([]byte("<h1>Hello World!!</h1>\n"))
	//	checkError(err)

	fmt.Println("<<<< end")
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("fatal: error %s\n", err.Error())
		os.Exit(1)
	}
}
