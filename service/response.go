package service

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func ReturnResponse(conn *net.TCPConn, method string, uri string, ver string, contentType string) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/www" + uri

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recover!:", err)

			_, err = conn.Write([]byte(ver + " 500 Internal Server Error\n"))
			_, err = conn.Write([]byte("Content-Type: text/html \n"))
			_, err = conn.Write([]byte("\n"))
			_, err = conn.Write([]byte("<html><head><h1>Internal Server Error</h1></head></html>"))

		}
	}()

	switch method {
	case "GET":
		if _, err := os.Stat(filePath); err != nil {
			checkError(err, ver, conn)
			body, err := ioutil.ReadFile(pwd + "/notfound.html")
			checkError(err, ver, conn)
			_, err = conn.Write([]byte(ver + " 404 Not Found\n"))
			checkError(err, ver, conn)

			// header
			_, err = conn.Write([]byte("Content-Type: " + contentType + "\n"))
			checkError(err, ver, conn)
			_, err = conn.Write([]byte("\n"))
			checkError(err, ver, conn)

			// body
			_, err = conn.Write([]byte(body))
			checkError(err, ver, conn)
			break
		}

		body, err := ioutil.ReadFile(filePath)

		// panic("Panic!")
		checkError(err, ver, conn)
		// status line
		_, err = conn.Write([]byte(ver + " 200 OK\n"))
		checkError(err, ver, conn)

		// header
		_, err = conn.Write([]byte("Content-Type: " + contentType + "\n"))
		checkError(err, ver, conn)
		_, err = conn.Write([]byte("\n"))
		checkError(err, ver, conn)

		// body
		_, err = conn.Write([]byte(body))
		checkError(err, ver, conn)
	case "HEAD":
		if _, err := os.Stat(filePath); err != nil {
			body, err := ioutil.ReadFile(pwd + "/notfound.html")
			checkError(err, ver, conn)
			_, err = conn.Write([]byte(ver + " 404 Not Found\n"))
			_, err = conn.Write([]byte("\n"))
			_, err = conn.Write([]byte(body))
			checkError(err, ver, conn)
			break
		}
		body, err := ioutil.ReadFile(filePath)
		checkError(err, ver, conn)
		// status line
		_, err = conn.Write([]byte(ver + " 200 OK\n"))
		checkError(err, ver, conn)
		_, err = conn.Write([]byte("\n"))
		checkError(err, ver, conn)
		_, err = conn.Write([]byte(body))
		checkError(err, ver, conn)
	default:
		_, err := conn.Write([]byte(ver + " 405 Method Not Allowed\n"))
		checkError(err, ver, conn)
		fmt.Println(method)
	}
}

func checkError(err error, ver string, conn *net.TCPConn) {
	if err != nil {
		fmt.Printf("fatal: error %s\n", err)
	}
}
