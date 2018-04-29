package service

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReturnResponse(method string, uri string, ver string, contentType string) (responseArray []string) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/www" + uri

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recover!:", err)

			responseArray = append(responseArray, ver+" 500 Internal Server Error\n")
			responseArray = append(responseArray, "Content-Type: text/html \n")
			responseArray = append(responseArray, "\n")
			responseArray = append(responseArray, "<html><head><h1>Internal Server Error</h1></head></html>")
		}
	}()

	switch method {
	case "GET":
		if _, err := os.Stat(filePath); err != nil {
			checkError(err)

			body, err := ioutil.ReadFile(pwd + "/www/notfound.html")
			checkError(err)

			responseArray = append(responseArray, ver+" 404 Not Found\n")
			responseArray = append(responseArray, "Content-Type: "+contentType+"\n")
			responseArray = append(responseArray, "\n")
			responseArray = append(responseArray, string(body))
			return responseArray
		}

		body, err := ioutil.ReadFile(filePath)
		// panic("Panic!")
		checkError(err)

		responseArray = append(responseArray, ver+" 200 OK\n")
		responseArray = append(responseArray, "Content-Type: "+contentType+"\n")
		responseArray = append(responseArray, "\n")
		responseArray = append(responseArray, string(body))
	case "HEAD":
		if _, err := os.Stat(filePath); err != nil {
			body, err := ioutil.ReadFile(pwd + "/notfound.html")
			checkError(err)

			responseArray = append(responseArray, ver+" 404 Not Found\n")
			responseArray = append(responseArray, "\n")
			responseArray = append(responseArray, string(body))
			return responseArray
		}
		body, err := ioutil.ReadFile(filePath)
		checkError(err)
		// status line
		responseArray = append(responseArray, ver+" 200 OK\n")
		responseArray = append(responseArray, "\n")
		responseArray = append(responseArray, string(body))
	default:
		responseArray = append(responseArray, ver+" 405 Method Not Allowed\n")
		responseArray = append(responseArray, "\n")
		responseArray = append(responseArray, "<html><head><h1>Internal Server Error</h1></head></html>")
		fmt.Println(method)
	}
	return responseArray
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("fatal: error %s\n", err)
	}
}
