package main

import "testing"

func TestRetrieveMethodRequestLine(t *testing.T) {
	if retrieveMethodRequestLine("GET /test HTTP1.0") != "GET" {
		t.Fatal("retrieveMethodRequestLine should be GET")
	}
}

func TestRetrieveURIRequestLinet(t *testing.T) {
	if retrieveURIRequestLine("GET /test HTTP1.0") != "/test" {
		t.Fatal("retrieveURIRequestLine should be /test")
	}
}

func TestRetrieveVersionRequestLine(t *testing.T) {
	if retrieveVersionRequestLine("GET /test HTTP1.0") != "HTTP1.0" {
		t.Fatal("retrieveVersionRequestLine should be HTTP1.0")
	}
}
