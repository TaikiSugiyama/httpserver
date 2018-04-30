package test

import (
	"fmt"
	"httpserver/server/service"
	"testing"
)

func TestReturnResponse_GET(t *testing.T) {
	actual := []string{"HTTP1.0 200 OK\n", "Content-Type: text/html\n", "\n", "<html></html>\n"}
	result := service.ReturnResponse("GET", "/test.html", "HTTP1.0", "text/html")
	fmt.Println(actual)
	fmt.Println(result)
	for i, v := range result {
		if actual[i] != v {
			t.Fatal("response should be " + result[i])
		}
	}
}

func TestReturnResponse_HEAD(t *testing.T) {
	actual := []string{"HTTP1.0 200 OK\n", "\n", "<html></html>\n"}
	result := service.ReturnResponse("HEAD", "/test.html", "HTTP1.0", "text/html")
	fmt.Println(actual)
	fmt.Println(result)
	for i, v := range result {
		if actual[i] != v {
			t.Fatal("response should be " + result[i])
		}
	}
}
