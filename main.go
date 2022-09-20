package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootHandler() {

}

func main() {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
