package service

import (
	"strings"
)

func RetrieveMethodRequestLine(line string) (method string) {
	split := strings.Split(line, " ")
	return split[0]
}

func RetrieveURIRequestLine(line string) (uri string) {
	split := strings.Split(line, " ")
	if split[1] == "/" {
		split[1] = "/index.html"
	}
	return split[1]
}

func RetrieveVersionRequestLine(line string) (ver string) {
	split := strings.Split(line, " ")
	return split[2]
}

func RetrieveTypeRequestLine(uri string) (extension string) {
	start := strings.LastIndex(uri, ".")
	r := []rune(uri)
	return string(r[start+1:])
}
