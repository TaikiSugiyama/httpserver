package handler

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("fatal: error %s\n", err.Error())
		os.Exit(1)
	}
}
