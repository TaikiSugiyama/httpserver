package test

import (
	"httpserver/service"
	"testing"
)

func TestRetrieveMethodRequestLine(t *testing.T) {
	if service.RetrieveMethodRequestLine("GET /test.html HTTP1.0") != "GET" {
		t.Fatal("retrieveMethodRequestLine should be GET")
	}
}

func TestRetrieveURIRequestLinet(t *testing.T) {
	if service.RetrieveURIRequestLine("GET /test.html HTTP1.0") != "/test.html" {
		t.Fatal("retrieveURIRequestLine should be /test")
	}
}

func TestRetrieveVersionRequestLine(t *testing.T) {
	if service.RetrieveVersionRequestLine("GET /test.html HTTP1.0") != "HTTP1.0" {
		t.Fatal("retrieveVersionRequestLine should be HTTP1.0")
	}
}

func TestRetrieveRequestLine(t *testing.T) {
	if service.RetrieveTypeRequestLine("/test.html") != "html" {
		t.Fatal("service.RetrieveTypeRequestLine should be html")
	}

	if service.RetrieveTypeRequestLine("/test.css") != "css" {
		t.Fatal("service.RetrieveTypeRequestLine should be css")
	}

	if service.RetrieveTypeRequestLine("/test.jpg") != "jpg" {
		t.Fatal("service.RetrieveTypeRequestLine should be jpg")
	}

	if service.RetrieveTypeRequestLine("/test.png") != "png" {
		t.Fatal("service.RetrieveTypeRequestLine should be png")
	}

	if service.RetrieveTypeRequestLine("/test.gif") != "gif" {
		t.Fatal("service.RetrieveTypeRequestLine should be gif")
	}
}
