package handler

import (
	"fmt"
	"httpserver/service"
	"net"
	"regexp"
)

func HandleListener(listener *net.TCPListener) error {
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

		go HandleConnection(conn)
	}
}

func readLine(conn *net.TCPConn) string {
	lineBuf := make([]byte, 1024)
	lineLen, err := conn.Read(lineBuf)
	checkError(err)

	line := string(lineBuf[:lineLen])

	// read確認用
	fmt.Println(line)
	return line
}

func HandleConnection(conn *net.TCPConn) {
	defer conn.Close()
	fmt.Println("start >>>")
	line := readLine(conn)

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

	method = service.RetrieveMethodRequestLine(requestArray[0])
	uri = service.RetrieveURIRequestLine(requestArray[0])
	contentType = typeMap[service.RetrieveTypeRequestLine(uri)]
	ver = service.RetrieveVersionRequestLine(requestArray[0])
	fmt.Println(contentType)

	service.ReturnResponse(conn, method, uri, ver, contentType)

	fmt.Println("<<<< end")
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("fatal: error %s\n", err.Error())
	}
}
