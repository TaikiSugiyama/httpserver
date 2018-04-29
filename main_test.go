package main

import (
	"testing"
)

func TestRetrieveMethodRequestLine(t *testing.T) {
	if retrieveMethodRequestLine("GET /test.html HTTP1.0") != "GET" {
		t.Fatal("retrieveMethodRequestLine should be GET")
	}
}

func TestRetrieveURIRequestLinet(t *testing.T) {
	if retrieveURIRequestLine("GET /test.html HTTP1.0") != "/test.html" {
		t.Fatal("retrieveURIRequestLine should be /test")
	}
}

func TestRetrieveVersionRequestLine(t *testing.T) {
	if retrieveVersionRequestLine("GET /test.html HTTP1.0") != "HTTP1.0" {
		t.Fatal("retrieveVersionRequestLine should be HTTP1.0")
	}
}

func TestRetrieveTypeRequestLine(t *testing.T) {
	if retrieveTypeRequestLine("/test.html") != ".html" {
		t.Fatal("retrieveTypeRequestLine should be .html")
	}

	if retrieveTypeRequestLine("/test.css") != ".css" {
		t.Fatal("retrieveTypeRequestLine should be .css")
	}
}
