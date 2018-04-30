package handler

import (
	"fmt"
	"httpserver/server/service"
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

	fmt.Println(requestArray[0])

	method = service.RetrieveMethodRequestLine(requestArray[0])
	uri = service.RetrieveURIRequestLine(requestArray[0])
	fmt.Println(uri)
	contentType = typeMap[service.RetrieveTypeRequestLine(uri)]
	ver = service.RetrieveVersionRequestLine(requestArray[0])
	fmt.Println(contentType)

	array := service.ReturnResponse(method, uri, ver, contentType)

	for _, v := range array {
		_, err := conn.Write([]byte(v))
		checkError(err)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Printf("fatal: error %s\n", err.Error())
	}
}
